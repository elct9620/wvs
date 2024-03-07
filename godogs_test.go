package wvs_test

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/elct9620/wvs"
)

var opts = godog.Options{
	Tags:   "~@wip",
	Format: "pretty",
	Paths:  []string{"features"},
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

type srvCtxKey struct{}

var (
	ErrServerNotFound = errors.New("server not found")
)

func beforeScenarioAppHook(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	app, err := wvs.InitializeTest()
	if err != nil {
		return ctx, err
	}

	srv := httptest.NewServer(app)
	return context.WithValue(ctx, srvCtxKey{}, srv), nil
}

func afterScenarioAppHook(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	srv, err := GetServer(ctx)
	if err != nil {
		return ctx, err
	}

	srv.Close()
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

	ctx.Step(`^the session id is "([^"]*)"$`, theSessionIdIs)
	ctx.Step(`^I make a (GET|POST|PUT|DELETE) request to "([^"]*)"$`, iMakeARequestTo)
	ctx.Step(`^the response body should be a valid JSON$`, theResponseBodyShouldBeAValidJson)
	ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
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
