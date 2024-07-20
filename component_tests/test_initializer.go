package component_tests

import (
	"component-tests/steps"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	s := steps.NewStep()
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, s.ISendrequestTo)
	ctx.Step(`^the response code should be (\d+)$`, s.TheResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, s.TheResponseShouldMatchJSON)
}
