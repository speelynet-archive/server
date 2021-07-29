package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"net/http/httptest"
	"testing"
)

var (
	ms *httptest.Server
)

func iCreateAMockServer() error {
	ms = httptest.NewServer(createRouter())
	return godog.ErrPending
}

func iCreateTheTemporaryFile(filepath string) error {
	return godog.ErrPending
}

func iGoToMockServer(url string) error {
	return godog.ErrPending
}

func iShouldSeeTheFileContent() error {
	return godog.ErrPending
}

func theContentOfTheTemporaryFileIs(content *godog.DocString) error {
	return godog.ErrPending
}

func CommonStepDefs(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a mock server$`, iCreateAMockServer)
	ctx.Step(`^I create the temporary file "([^"]*)"$`, iCreateTheTemporaryFile)
	ctx.Step(`^I go to "mock.server/([^"]*)"$`, iGoToMockServer)
	ctx.Step(`^I should see the file content$`, iShouldSeeTheFileContent)
	ctx.Step(`^the content of the temporary file is$`, theContentOfTheTemporaryFileIs)

	ctx.AfterScenario(func(s *godog.Scenario, e error) {
		if ms != nil {
			ms.Close()
		}
		ms = nil
	})
}

var opts = godog.Options{
	Format:    "pretty",
	Randomize: -1,
}

func RunFeature(t *testing.T, name string, TestSuiteInitializer func(ctx *godog.TestSuiteContext), ScenarioInitializer func(ctx *godog.ScenarioContext)) {
	o := opts
	o.Paths = []string{fmt.Sprintf("features/%s.feature", name)}

	if status := (godog.TestSuite{
		Name:                 name,
		TestSuiteInitializer: TestSuiteInitializer,
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			CommonStepDefs(ctx)
			ScenarioInitializer(ctx)
		},
		Options: &o,
	}).Run(); status != 0 {
		t.Errorf("godog exited with non-zero code '%d'", status)
	}
}
