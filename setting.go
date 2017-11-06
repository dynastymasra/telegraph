package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type (
	GetUpdateCall struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	SetWebHookCall struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

// GetUpdate prepare request telegram get update
func (client *Client) GetUpdate() *GetUpdateCall {
	url := client.baseURL + fmt.Sprintf(EndpointGetUpdate, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &GetUpdateCall{
		Client:  client,
		Request: request,
	}
}

// QueryOffset add request with query params offset
func (call *GetUpdateCall) QueryOffset(offset int) *GetUpdateCall {
	return &GetUpdateCall{
		Client:  call.Client,
		Request: call.Request.Query(fmt.Sprintf("offset=%v", offset)),
	}
}

// QueryLimit add request with query params limit
func (call *GetUpdateCall) QueryLimit(limit int) *GetUpdateCall {
	return &GetUpdateCall{
		Client:  call.Client,
		Request: call.Request.Query(fmt.Sprintf("limit=%v", limit)),
	}
}

// QueryTimeout add request with query params timeout
func (call *GetUpdateCall) QueryTimeout(timeout int) *GetUpdateCall {
	return &GetUpdateCall{
		Client:  call.Client,
		Request: call.Request.Query(fmt.Sprintf("timeout=%v", timeout)),
	}
}

// QueryAllowedUpdate add request with query params allowed updates
func (call *GetUpdateCall) QueryAllowedUpdate(updates ...string) *GetUpdateCall {
	return &GetUpdateCall{
		Client:  call.Client,
		Request: call.Request.Query(fmt.Sprintf("allowed_updates=%v", updates)),
	}
}

// Commit make request get update telegram
//func (call *GetUpdateCall) Commit() (*http.Response, []byte, error) {
//	prepareRequest := PrepareRequest{
//		Client:  call.Client,
//		Request: call.Request,
//	}
//	return prepareRequest.Commit()
//}

// SetWebHook prepare request telegram api set web hook
func (client *Client) SetWebHook(webHook string) *SetWebHookCall {
	body := JSON{
		"url": webHook,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetWebHook, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &SetWebHookCall{
		Client:  client,
		Request: request,
	}
}

// Certificate set telegram web hook with certificate
func (call *SetWebHookCall) Certificate(path string) *SetWebHookCall {
	return &SetWebHookCall{
		Client:  call.Client,
		Request: call.Request.Type(gorequest.TypeMultipart).SendFile(path, "", "certificate"),
	}
}

// MaxConnection set telegram web hook with max connection, default 40
func (call *SetWebHookCall) MaxConnection(conn int) *SetWebHookCall {
	body := JSON{
		"max_connections": conn,
	}

	return &SetWebHookCall{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// AllowedUpdate set telegram web hook with allowed updates
func (call *SetWebHookCall) AllowedUpdate(allowed ...string) *SetWebHookCall {
	body := JSON{
		"allowed_updates": allowed,
	}

	return &SetWebHookCall{
		Client:  call.Client,
		Request: call.Request.Send(body),
	}
}

// Commit make request set web hook telegram
//func (call *SetWebHookCall) Commit() (*http.Response, []byte, error) {
//	prepareRequest := PrepareRequest{
//		Client:  call.Client,
//		Request: call.Request,
//	}
//
//	return prepareRequest.Commit()
//}

// DeleteWebHook request to delete telegram web hook
func (client *Client) DeleteWebHook() *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(EndpointDeleteWebHook, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		Client:  client,
		Request: request,
	}
}

// GetWebHookInfo request get info web hook telegram
func (client *Client) GetWebHookInfo() *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(EndpointGetWebHookInfo, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		Client:  client,
		Request: request,
	}
}
