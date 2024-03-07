package wvs_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/gorilla/websocket"
)

type wsCtxKey struct{}

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

func theWebsocketEventIsReceived(ctx context.Context, expectedJson *godog.DocString) error {
	ws, err := GetWebSocket(ctx)
	if err != nil {
		return err
	}

	var expected any
	if err := json.Unmarshal([]byte(expectedJson.Content), &expected); err != nil {
		return err
	}

	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	events := make(chan any)

	go func() {
		for {
			var event any
			if err := ws.ReadJSON(&event); err != nil {
				cancel()
				return
			}

			events <- event
		}
	}()

	for {
		select {
		case <-timeout.Done():
			return ErrWaitTimeout
		case actual := <-events:
			if reflect.DeepEqual(actual, expected) {
				return nil
			}
		}
	}
}
