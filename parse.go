package telegraph

import (
	"net/http"

	"fmt"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// VoidResponse struct to handle request and response telegram api
	VoidResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// ErrorResponse struct parse error response from telegram
	ErrorResponse struct {
		OK          bool   `json:"ok"`
		ErrorCode   int    `json:"error_code,omitempty"`
		Description string `json:"description,omitempty"`
	}
)

// MakeHTTPResponse create mock http response if request to API is error internal
func MakeHTTPResponse(agent *gorequest.SuperAgent) *http.Response {
	request, err := agent.MakeRequest()
	if err != nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}
	}

	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     request.Header,
		Request:    request,
	}
}

// Commit request to telegram api
func (call *VoidResponse) Commit() (int, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result bool `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = call.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, call.Client.expBackOff); err != nil {
		return http.StatusInternalServerError, err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return res.StatusCode, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return res.StatusCode, nil
}

// Download response from telegram API cannot used with Commit
func (call *VoidResponse) Download() (*http.Response, []byte, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = call.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, call.Client.expBackOff); err != nil {
		return nil, nil, err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return nil, nil, fmt.Errorf("%v", body)
	}
	return res, body, nil
}
