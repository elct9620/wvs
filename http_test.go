package wvs_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/elct9620/wvs/pkg/session"
	"github.com/google/go-cmp/cmp"
)

type httpResCtxKey struct{}
type sessionIdCtxKey struct{}

const DefaultSessionKey = "1234567890123456"

var (
	ErrHttpResponseNotFound = errors.New("http response not found in context")
)

func GetResponse(ctx context.Context) (*http.Response, error) {
	if res, ok := ctx.Value(httpResCtxKey{}).(*http.Response); ok {
		return res, nil
	}

	return nil, ErrHttpResponseNotFound
}

func newRequest(ctx context.Context, method, target string, body io.Reader) (*http.Request, error) {
	srv, err := GetServer(ctx)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", srv.URL, target)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	if sessionId, ok := ctx.Value(sessionIdCtxKey{}).(string); ok {
		req.AddCookie(&http.Cookie{
			Name:  session.DefaultCookieName,
			Value: sessionId,
		})
	}

	return req, nil
}

func theSessionIdIs(ctx context.Context, id string) (context.Context, error) {
	encryptedSessionId, err := session.Encrypt([]byte(id), []byte(DefaultSessionKey))
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, sessionIdCtxKey{}, encryptedSessionId), nil
}

func iMakeARequestTo(ctx context.Context, method string, target string) (context.Context, error) {
	req, err := newRequest(ctx, method, target, nil)
	if err != nil {
		return ctx, err
	}

	srv, err := GetServer(ctx)
	if err != nil {
		return ctx, err
	}

	client := srv.Client()
	res, err := client.Do(req)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, httpResCtxKey{}, res), nil
}

func iMakeARequestToWithBody(ctx context.Context, method string, target string, body *godog.DocString) (context.Context, error) {
	req, err := newRequest(ctx, method, target, bytes.NewBufferString(body.Content))
	if err != nil {
		return ctx, err
	}

	srv, err := GetServer(ctx)
	if err != nil {
		return ctx, err
	}

	client := srv.Client()
	res, err := client.Do(req)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, httpResCtxKey{}, res), nil
}

func theResponseBodyShouldBeAValidJson(ctx context.Context, expectedJson *godog.DocString) error {
	var expected, actual any

	res, err := GetResponse(ctx)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(expectedJson.Content), &expected); err != nil {
		return err
	}

	actualBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(actualBody, &actual); err != nil {
		return fmt.Errorf("response body is not a valid JSON: %s", actualBody)
	}

	if diff := cmp.Diff(expected, actual); diff != "" {
		return fmt.Errorf("the response mismatch (-want, +got):\n%s", diff)
	}

	return nil
}

func theResponseStatusCodeShouldBe(ctx context.Context, code int) error {
	res, err := GetResponse(ctx)
	if err != nil {
		return err
	}

	if res.StatusCode != code {
		return fmt.Errorf("expected status code %d, but got %d", code, res.StatusCode)
	}

	return nil
}
