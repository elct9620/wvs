package wvs_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/elct9620/wvs/internal/app"
	"github.com/elct9620/wvs/pkg/session"
	"github.com/google/go-cmp/cmp"
)

type httpResCtxKey struct{}
type sessionIdCtxKey struct{}

const DefaultSessionKey = "1234567890123456"

var (
	ErrHttpResponseNotFound = errors.New("http response not found in context")
)

func newRequest(ctx context.Context, method, target string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, target, body)

	if sessionId, ok := ctx.Value(sessionIdCtxKey{}).(string); ok {
		req.AddCookie(&http.Cookie{
			Name:  session.DefaultCookieName,
			Value: sessionId,
		})
	}

	return req
}

func theSessionIdIs(ctx context.Context, id string) (context.Context, error) {
	encryptedSessionId, err := session.Encrypt([]byte(id), []byte(DefaultSessionKey))
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, sessionIdCtxKey{}, encryptedSessionId), nil
}

func iMakeARequestTo(ctx context.Context, method string, target string) (context.Context, error) {
	req := newRequest(ctx, method, target, nil)
	res := httptest.NewRecorder()

	app, ok := ctx.Value(appCtxKey{}).(*app.Application)
	if !ok {
		return nil, ErrAppNotFound
	}

	app.ServeHTTP(res, req)
	return context.WithValue(ctx, httpResCtxKey{}, res), nil
}

func theResponseBodyShouldBeAValidJson(ctx context.Context, expectedJson *godog.DocString) error {
	var expected, actual any

	res, ok := ctx.Value(httpResCtxKey{}).(*httptest.ResponseRecorder)
	if !ok {
		return ErrHttpResponseNotFound
	}

	if err := json.Unmarshal([]byte(expectedJson.Content), &expected); err != nil {
		return err
	}

	actualBody := res.Body.Bytes()
	if err := json.Unmarshal(actualBody, &actual); err != nil {
		return fmt.Errorf("response body is not a valid JSON: %s", actualBody)
	}

	if diff := cmp.Diff(expected, actual); diff != "" {
		return fmt.Errorf("the response mismatch (-want, +got):\n%s", diff)
	}

	return nil
}

func theResponseStatusCodeShouldBe(ctx context.Context, code int) error {
	res, ok := ctx.Value(httpResCtxKey{}).(*httptest.ResponseRecorder)
	if !ok {
		return ErrHttpResponseNotFound
	}

	if res.Code != code {
		return fmt.Errorf("expected status code %d, but got %d", code, res.Code)
	}

	return nil
}
