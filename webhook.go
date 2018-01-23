package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// WebHookInfoResponse struct to handle request and response telegram api
	WebHookInfoResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

// GetWebHookInfo Use this method to get current webhook status. Requires no parameters.
// On success, returns a WebhookInfo object. If the bot is using getUpdates,
// will return an object with the url field empty.
func (client *Client) GetWebHookInfo() *WebHookInfoResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetWebHookInfo, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &WebHookInfoResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (info *WebHookInfoResponse) Commit() (*WebhookInfo, *http.Response, error) {
	var errs []error
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *WebhookInfo `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, _, errs = info.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, info.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(info.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}
