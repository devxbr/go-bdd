package main

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

var calc Calculator

func iHaveACalculator() error {
	calc = Calculator{}
	return nil
}

func iAddAnd(a, b int) error {
	calc.add(a, b)
	return nil
}

func theResultShouldBe(expected int) error {
	if calc.result != expected {
		return fmt.Errorf("expected result to be %d but got %d", expected, calc.result)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have a calculator$`, iHaveACalculator)
	ctx.Step(`^I add (\d+) and (\d+)$`, iAddAnd)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
