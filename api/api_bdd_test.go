package main

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestApiBdd(t *testing.T) {
	opts := godog.Options{
		Format:    "pretty", // Use "pretty" for human-readable output
		Paths:     []string{"api_bdd.feature"},
		Randomize: 0, // Randomize scenario execution order
	}

	suite := godog.TestSuite{
		Name:                "godog",
		ScenarioInitializer: func(ctx *godog.ScenarioContext) { InitializeScenario(ctx) },
		Options:             &opts,
	}

	if suite.Run() != 0 {
		t.Fail()
	}
}
