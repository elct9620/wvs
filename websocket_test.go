package wvs_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type wsCtxKey struct{}

var (
	MaxWebsocketWaitTimeout = 5 * time.Second
)

var (
	ErrNoWebsocket = errors.New("no websocket in context")
	ErrWaitTimeout = errors.New("wait timeout")
)

func GetWebSocket(ctx context.Context) (*websocket.Conn, error) {
	ws, ok := ctx.Value(wsCtxKey{}).(*websocket.Conn)
	if !ok {
		return nil, ErrNoWebsocket
	}
	return ws, nil
}

func readWebsocketEvents(ctx context.Context) (chan any, error) {
	ws, err := GetWebSocket(ctx)
	if err != nil {
		return nil, err
	}

	events := make(chan any)

	go func() {
		for {
			var event any
			if err := ws.ReadJSON(&event); err != nil {
				return
			}

			events <- event
		}
	}()

	return events, nil
}

func connectToTheWebsocket(ctx context.Context) (context.Context, error) {
	srv, err := GetServer(ctx)
	if err != nil {
		return ctx, err
	}

	url := fmt.Sprintf("ws://%s/ws", strings.TrimPrefix(srv.URL, "http://"))
	ws, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return ctx, err
	}

	newCtx := context.WithValue(ctx, wsCtxKey{}, ws)
	return context.WithValue(newCtx, httpResCtxKey{}, res), nil
}

func theWebsocketEventIsReceived(ctx context.Context, eventType string) error {
	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	events, err := readWebsocketEvents(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case <-timeout.Done():
			return ErrWaitTimeout
		case actual := <-events:
			actualValue, ok := actual.(map[string]interface{})
			if !ok {
				continue
			}

			if actualValue["type"] == eventType {
				return nil
			}
		}
	}
}
