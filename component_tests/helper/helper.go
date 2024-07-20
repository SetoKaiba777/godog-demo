package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/cucumber/godog"
)

type Helper struct {
	client *http.Client
	resp   *http.Response
}

func NewHelper() *Helper {
	return &Helper{client: http.DefaultClient}
}

func (h *Helper) SendRequest(method, endpoint string) error {
	req, err := http.NewRequest(method, "http://localhost:8080"+endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	h.resp = resp
	return nil
}

func (h *Helper) CompareStatusCode(code int) error {
	if code != h.resp.StatusCode {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, h.resp.StatusCode)
	}
	return nil
}

func (h *Helper) CompareBody(body *godog.DocString) error {
	var expected, actual interface{}

	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return fmt.Errorf("error in unmarshal body: %v", err)
	}

	bodyBytes, err := io.ReadAll(h.resp.Body)
	if err != nil {
		return fmt.Errorf("error in read body: %v", err)
	}

	if err = json.Unmarshal(bodyBytes, &actual); err != nil {
		return fmt.Errorf("error in unmarshal body: %v", err)
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}
