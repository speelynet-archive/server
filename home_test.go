package main_test

import (
	"github.com/cucumber/godog"
	"testing"
)

func InitializeHomeSuite(_ *godog.TestSuiteContext)   {}
func InitializeHomeScenario(_ *godog.ScenarioContext) {}

func TestHome(t *testing.T) {
	RunFeature(t, "home", InitializeHomeSuite, InitializeHomeScenario)
}
