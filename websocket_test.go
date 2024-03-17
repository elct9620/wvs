package wvs_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"

	"github.com/elct9620/wvs/pkg/session"
	"github.com/gorilla/websocket"
	"github.com/jmespath/go-jmespath"
)

type wsCtxKey struct{}

var (
	MaxWebsocketWaitTimeout = 5 * time.Second
)

var (
	ErrNoWebsocket = errors.New("no websocket in context")
	ErrWaitTimeout = errors.New("wait timeout")
)

type WebSocketClient struct {
	mux sync.RWMutex
	*websocket.Conn
	events []any
}

func NewWebSocketClient(ws *websocket.Conn) *WebSocketClient {
	return &WebSocketClient{
		Conn:   ws,
		events: make([]any, 0),
	}
}

func (c *WebSocketClient) Events() []any {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.events
}

func (c *WebSocketClient) ReadEvents() {
	for {
		var event any
		if err := c.ReadJSON(&event); err != nil {
			break
		}

		c.mux.Lock()
		c.events = append(c.events, event)
		c.mux.Unlock()
	}
}

func GetWebSocket(ctx context.Context) (*WebSocketClient, error) {
	ws, ok := ctx.Value(wsCtxKey{}).(*WebSocketClient)
	if !ok {
		return nil, ErrNoWebsocket
	}

	return ws, nil
}

func connectToTheWebsocket(ctx context.Context) (context.Context, error) {
	srv, err := GetServer(ctx)
	if err != nil {
		return ctx, err
	}

	url, err := url.Parse(srv.URL)
	if err != nil {
		return ctx, err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return ctx, err
	}

	dailer := websocket.Dialer{Jar: jar}

	if sessionId, ok := ctx.Value(sessionIdCtxKey{}).(string); ok {
		dailer.Jar.SetCookies(url, []*http.Cookie{
			{
				Name:  session.DefaultCookieName,
				Value: sessionId,
			},
		})
	}

	wsUrl := fmt.Sprintf("ws://%s/ws", url.Host)
	ws, res, err := dailer.Dial(wsUrl, nil)
	if err != nil {
		if errors.Is(err, websocket.ErrBadHandshake) {
			return context.WithValue(ctx, httpResCtxKey{}, res), nil
		}

		return ctx, err
	}

	client := NewWebSocketClient(ws)
	go client.ReadEvents()

	newCtx := context.WithValue(ctx, wsCtxKey{}, client)
	return context.WithValue(newCtx, httpResCtxKey{}, res), nil
}

func theWebsocketEventIsReceived(ctx context.Context, eventType string) error {
	ws, err := GetWebSocket(ctx)
	if err != nil {
		return err
	}

	timeout := time.After(MaxWebsocketWaitTimeout)
	for {
		select {
		case <-timeout:
			return ErrWaitTimeout
		default:
			for _, event := range ws.Events() {
				event, ok := event.(map[string]interface{})
				if !ok {
					continue
				}

				if event["type"] == eventType {
					return nil
				}
			}
		}
	}
}

func theWebsocketEventHasWithValue(ctx context.Context, eventType, path, value string) error {
	ws, err := GetWebSocket(ctx)
	if err != nil {
		return err
	}

	timeout := time.After(MaxWebsocketWaitTimeout)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("event with path %s and value %s not found in: \n%+v", path, value, ws.Events())
		default:
			for _, event := range ws.Events() {
				event, ok := event.(map[string]interface{})
				if !ok {
					continue
				}

				if event["type"] != eventType {
					continue
				}

				res, err := jmespath.Search(path, event)
				if err != nil {
					return err
				}

				if res == value {
					return nil
				}
			}
		}
	}
}
