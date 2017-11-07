package telegraph

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// WebhookInfo Contains information about the current status of a webhook.
	WebhookInfo struct {
		URL                  string   `json:"url"`
		HasCustomCertificate bool     `json:"has_custom_certificate"`
		PendingUpdateCount   int      `json:"pending_update_count"`
		LastErrorDate        int64    `json:"last_error_date,omitempty"`
		LastErrorMessage     string   `json:"last_error_message,omitempty"`
		MaxConnections       int      `json:"max_connections,omitempty"`
		AllowedUpdates       []string `json:"allowed_updates,omitempty"`
	}

	InfoResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
Use this method to specify a url and receive incoming updates via an outgoing webhook.
Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
containing a JSON-serialized Update. In case of an unsuccessful request,
we will give up after a reasonable amount of attempts. Returns true.
*/
func (client *Client) SetWebHook(webHook string) *VoidResponse {
	body := JSON{
		"url": webHook,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetWebHook, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetCertificate Upload your public key certificate so that the root certificate in use can be checked.
// See our self-signed guide for details.
func (call *VoidResponse) SetCertificate(path string) *VoidResponse {
	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Type(gorequest.TypeMultipart).SendFile(path, "", "certificate"),
	}
}

// SetMaxConnection Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery,
// 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server,
// and higher values to increase your bot’s throughput.
func (call *VoidResponse) SetMaxConnection(conn int) *VoidResponse {
	body := JSON{
		"max_connections": conn,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

/*
SetAllowedUpdates List the types of updates you want your bot to receive.
For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types.
See Update for a complete list of available update types.
Specify an empty list to receive all updates regardless of type (default).
If not specified, the previous setting will be used.
*/
func (call *VoidResponse) SetAllowedUpdates(allowed ...string) *VoidResponse {
	body := JSON{
		"allowed_updates": allowed,
	}

	return &VoidResponse{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// DeleteWebHook Use this method to remove webhook integration if you decide to switch back to getUpdates.
// Returns True on success. Requires no parameters.
func (client *Client) DeleteWebHook() *VoidResponse {
	url := client.baseURL + fmt.Sprintf(EndpointDeleteWebHook, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// GetWebHookInfo Use this method to get current webhook status. Requires no parameters.
// On success, returns a WebhookInfo object. If the bot is using getUpdates,
// will return an object with the url field empty.
func (client *Client) GetWebHookInfo() *InfoResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetWebHookInfo, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &InfoResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (info *InfoResponse) Commit() (*WebhookInfo, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = info.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, info.Client.expBackOff); err != nil {
		return nil, &http.Response{StatusCode: http.StatusInternalServerError}, err
	}
	return parseWebHookInfo(res, body)
}

func parseWebHookInfo(res *http.Response, body []byte) (*WebhookInfo, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *WebhookInfo `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}
	return model.Result, res, nil
}
