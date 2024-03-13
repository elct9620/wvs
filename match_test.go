package wvs_test

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cucumber/godog"
)

func thereHaveSomeMatch(ctx context.Context, payload *godog.DocString) error {
	req, err := newRequest(ctx, http.MethodPost, "/testability/matches", bytes.NewBufferString(payload.Content))
	if err != nil {
		return err
	}

	srv, err := GetServer(ctx)
	if err != nil {
		return err
	}

	client := srv.Client()
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
