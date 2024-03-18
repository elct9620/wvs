package wvs_test

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/elct9620/wvs"
	"github.com/elct9620/wvs/internal/app"
)

var opts = godog.Options{
	Tags:   "~@wip",
	Format: "pretty",
	Paths:  []string{"features"},
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

type appCtxKey struct{}
type srvCtxKey struct{}

var (
	ErrServerNotFound = errors.New("server not found")
)

func GetApp(ctx context.Context) (*app.Application, error) {
	if app, ok := ctx.Value(appCtxKey{}).(*app.Application); ok {
		return app, nil
	}

	return nil, errors.New("app not found")
}

func beforeScenarioAppHook(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	app, err := wvs.InitializeTest()
	if err != nil {
		return ctx, err
	}

	go app.Event.Run(ctx) // nolint:errcheck

	srv := httptest.NewServer(app)

	ctx = context.WithValue(ctx, appCtxKey{}, app)
	return context.WithValue(ctx, srvCtxKey{}, srv), nil
}

func afterScenarioAppHook(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	if srv, err := GetServer(ctx); err == nil {
		srv.Close()
	}

	if app, err := GetApp(ctx); err == nil {
		app.Event.Close() // nolint:errcheck
	}

	if ws, err := GetWebSocket(ctx); err == nil {
		ws.Close()
	}

	return ctx, nil
}

func GetServer(ctx context.Context) (*httptest.Server, error) {
	if srv, ok := ctx.Value(srvCtxKey{}).(*httptest.Server); ok {
		return srv, nil
	}

	return nil, ErrServerNotFound
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(beforeScenarioAppHook)
	ctx.After(afterScenarioAppHook)

	ctx.Step(`^the session id is "([^"]*)"$`, theSessionIdIs)
	ctx.Step(`^I make a (GET|POST|PUT|DELETE) request to "([^"]*)"$`, iMakeARequestTo)
	ctx.Step(`^I make a (GET|POST|PUT|DELETE) request to "([^"]*)" with body$`, iMakeARequestToWithBody)
	ctx.Step(`^the response body should be a valid JSON$`, theResponseBodyShouldBeAValidJson)
	ctx.Step(`^the response JSON should has "([^"]*)"$`, theResponseJsonShouldHas)
	ctx.Step(`^the response JSON should has "([^"]*)" with value "([^"]*)"$`, theResponseJSONShouldHasWithValue)
	ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)

	ctx.Step(`^connect to the websocket$`, connectToTheWebsocket)
	ctx.Step(`^the websocket event "([^"]*)" is received$`, theWebsocketEventIsReceived)
	ctx.Step(`^the websocket event "([^"]*)" has "([^"]*)" with value "([^"]*)"$`, theWebsocketEventHasWithValue)

	ctx.Step(`^there have some match$`, thereHaveSomeMatch)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                "walus-vs-slime",
		ScenarioInitializer: InitializeScenario,
		Options:             &o,
	}.Run()

	if status != 0 {
		t.Errorf("godog failed with status: %d", status)
	}
}
