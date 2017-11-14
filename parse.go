package telegraph

import (
	"net/http"

	"encoding/json"
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

	operation := func() error {
		res, body, errs = call.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, call.Client.expBackOff); err != nil {
		return http.StatusInternalServerError, err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return parseErrorResponse(res, body)
	}

	return res.StatusCode, nil
}

// Download response from telegram API
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
		_, err := parseErrorResponse(res, body)
		return nil, nil, err
	}
	return res, body, nil
}

func parseErrorResponse(res *http.Response, body []byte) (int, error) {
	var response ErrorResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return res.StatusCode, err
	}
	return res.StatusCode, fmt.Errorf(response.Description)
}
