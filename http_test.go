package wvs_test

import (
	"context"
	"errors"
	"fmt"
	"net/http/httptest"

	"github.com/elct9620/wvs/app"
)

type httpResCtxKey struct{}

var (
	ErrHttpResponseNotFound = errors.New("http response not found in context")
)

func iMakeARequestTo(ctx context.Context, method string, path string) (context.Context, error) {
	req := httptest.NewRequest(method, path, nil)
	res := httptest.NewRecorder()

	app, ok := ctx.Value(appCtxKey{}).(*app.Application)
	if !ok {
		return nil, ErrAppNotFound
	}

	app.ServeHTTP(res, req)
	return context.WithValue(ctx, httpResCtxKey{}, res), nil
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
