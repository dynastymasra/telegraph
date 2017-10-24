package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// GetUpdate prepare request telegram get update
func (client *Client) GetUpdate() *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(EndpointGetUpdate, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		client:  client,
		request: request,
	}
}

// QueryOffset add request with query params offset
func (call *PrepareRequest) QueryOffset(offset int) *PrepareRequest {
	return &PrepareRequest{
		client:  call.client,
		request: call.request.Query(fmt.Sprintf("offset=%v", offset)),
	}
}

// QueryLimit add request with query params limit
func (call *PrepareRequest) QueryLimit(limit int) *PrepareRequest {
	return &PrepareRequest{
		client:  call.client,
		request: call.request.Query(fmt.Sprintf("limit=%v", limit)),
	}
}

// QueryTimeout add request with query params timeout
func (call *PrepareRequest) QueryTimeout(timeout int) *PrepareRequest {
	return &PrepareRequest{
		client:  call.client,
		request: call.request.Query(fmt.Sprintf("timeout=%v", timeout)),
	}
}

// QueryAllowedUpdate add request with query params allowed updates
func (call *PrepareRequest) QueryAllowedUpdate(updates ...string) *PrepareRequest {
	return &PrepareRequest{
		client:  call.client,
		request: call.request.Query(fmt.Sprintf("allowed_updates=%v", updates)),
	}
}

// SetWebHook prepare request telegram api set web hook
func (client *Client) SetWebHook(webHook string) *PrepareRequest {
	body := JSON{
		"url": webHook,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetWebHook, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &PrepareRequest{
		client:  client,
		request: request,
	}
}

// Certificate set telegram web hook with certificate
func (call *PrepareRequest) Certificate(path string) *PrepareRequest {
	return &PrepareRequest{
		client:  call.client,
		request: call.request.Type(gorequest.TypeMultipart).SendFile(path, "", "certificate"),
	}
}

// MaxConnection set telegram web hook with max connection, default 40
func (call *PrepareRequest) MaxConnection(conn int) *PrepareRequest {
	body := JSON{
		"max_connections": conn,
	}

	return &PrepareRequest{
		client:  call.client,
		request: call.request.Send(body),
	}
}

// AllowedUpdate set telegram web hook with allowed updates
func (call *PrepareRequest) AllowedUpdate(allowed ...string) *PrepareRequest {
	body := JSON{
		"allowed_updates": allowed,
	}

	return &PrepareRequest{
		client:  call.client,
		request: call.request.Send(body),
	}
}

// DeleteWebHook request to delete telegram web hook
func (client *Client) DeleteWebHook() *PrepareRequest {
	url := BaseURL + fmt.Sprintf(EndpointDeleteWebHook, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		client:  client,
		request: request,
	}
}

// GetWebHookInfo request get info web hook telegram
func (client *Client) GetWebHookInfo() *PrepareRequest {
	url := BaseURL + fmt.Sprintf(EndpointGetWebHookInfo, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		client:  client,
		request: request,
	}
}
