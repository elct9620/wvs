package wvs_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/elct9620/wvs"
	"github.com/spf13/pflag"
)

var opts = godog.Options{}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

type appCtxKey struct{}

var (
	ErrAppNotFound = errors.New("app not found in context")
)

func beforeScenarioAppHook(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	app, err := wvs.InitializeTest()
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, appCtxKey{}, app), nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(beforeScenarioAppHook)

	ctx.Step(`^I make a (GET|POST|PUT|DELETE) request to "([^"]*)"$`, iMakeARequestTo)
	ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                "walus-vs-slime",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
