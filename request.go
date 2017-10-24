package telegraph

import (
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	JSON map[string]interface{}

	PrepareRequest struct {
		client  *Client
		request *gorequest.SuperAgent
	}
)

// Commit request to telegram api
func (call *PrepareRequest) Commit() (*http.Response, []byte, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = call.request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, call.client.expBackOff); err != nil {
		return nil, nil, err
	}
	return res, body, nil
}
