package wvs_test

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
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
