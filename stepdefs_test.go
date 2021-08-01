package main_test

import (
	"fmt"
	"github.com/cucumber/godog"
	. "github.com/speelynet/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
)

var (
	ms          *httptest.Server
	currentFile *os.File
	fileContent string
	response    string
)

func iCreateAMockServer() error {
	ms = httptest.NewServer(CreateRouter())
	return nil
}

func iCreateTheTemporaryFile(filepath string) error {
	dir := path.Dir(filepath)

	if e := os.MkdirAll(dir, os.ModePerm); e != nil {
		return e
	} else if currentFile, e = os.Create(filepath); e != nil {
		return e
	}

	return nil
}

func iGoToMockServer(url string) error {
	if r, e := http.Get(ms.URL + url); e != nil {
		return e
	} else if b, e := ioutil.ReadAll(r.Body); e != nil {
		return e
	} else {
		response = string(b)
	}
	return nil
}

func iShouldSeeTheFileContent() error {
	indent := func(s string) string {
		return "    " + strings.ReplaceAll(s, "\n", "\n    ")
	}

	if response != fileContent {
		return fmt.Errorf(indent(fmt.Sprintf("\nexpected:\n%s\ngot:\n%s\n", indent(fileContent), indent(response))))
	}
	return nil
}

func theContentOfTheTemporaryFileIs(content *godog.DocString) error {
	if _, e := currentFile.WriteString(content.Content); e != nil {
		return e
	}
	fileContent = content.Content

	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		_ = os.Mkdir("testdata", os.ModePerm)
		_ = os.Chdir("testdata")
	})
	ctx.AfterSuite(func() {
		_ = os.Chdir("..")
		_ = os.RemoveAll("testdata")
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a mock server$`, iCreateAMockServer)
	ctx.Step(`^I create the temporary file "([^"]*)"$`, iCreateTheTemporaryFile)
	ctx.Step(`^I go to "mock.server([^"]*)"$`, iGoToMockServer)
	ctx.Step(`^I should see the file content$`, iShouldSeeTheFileContent)
	ctx.Step(`^the content of the temporary file is$`, theContentOfTheTemporaryFileIs)

	ctx.AfterScenario(func(s *godog.Scenario, e error) {
		if ms != nil {
			ms.Close()
		}
		ms = nil

		if currentFile != nil {
			_ = currentFile.Close()
			_ = os.Remove(currentFile.Name())
		}
		currentFile = nil

		fileContent = ""

		_ = os.RemoveAll("*")
	})
}
