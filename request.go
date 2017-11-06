package telegraph

import (
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	JSON map[string]interface{}

	PrepareRequest struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

// Commit request to telegram api
func (call *PrepareRequest) Commit() *Response {
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
		return &Response{
			Response: nil,
			Body:     nil,
			Err:      err,
		}
	}
	return &Response{
		Response: res,
		Body:     body,
		Err:      nil,
	}
}
