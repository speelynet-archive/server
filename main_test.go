package main_test

import (
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
	"os"
	"strings"
	"testing"
)

var (
	opts = godog.Options{
		Format:    "pretty",
		Randomize: -1,
	}
	osargs     []string
	godogflags []string
)

func init() {
	godog.BindCommandLineFlags("godog.", &opts)

	for _, v := range os.Args[1:] {
		var ptr *[]string
		if strings.HasPrefix(v, "--godog.") {
			ptr = &godogflags
		} else {
			ptr = &osargs
		}
		*ptr = append(*ptr, v)
	}

	os.Args = osargs
}

func TestMain(m *testing.M) {
	_ = pflag.CommandLine.Parse(godogflags)
	_ = flag.CommandLine.Parse(osargs)
	os.Exit(m.Run())
}

func RunFeature(t *testing.T, name string) {
	o := opts
	o.Paths = []string{fmt.Sprintf("features/%s.feature", name)}

	if status := (godog.TestSuite{
		Name:                 name,
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &o,
	}).Run(); status != 0 {
		t.Errorf("godog exited with non-zero code '%d'", status)
	}
}
