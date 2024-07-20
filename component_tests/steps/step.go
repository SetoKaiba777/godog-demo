package steps

import (
	"component-tests/helper"

	"github.com/cucumber/godog"
)

type Step struct {
	h *helper.Helper
}

func NewStep() Step {
	return Step{ h : helper.NewHelper()}
}

func (s *Step) ISendrequestTo(method, endpoint string) error {
	return s.h.SendRequest(method,endpoint)
}

func (s *Step) TheResponseCodeShouldBe(code int) error {
	return s.h.CompareStatusCode(code)
}

func (s *Step) TheResponseShouldMatchJSON(body *godog.DocString) (err error) {
	return s.h.CompareBody(body)
}
