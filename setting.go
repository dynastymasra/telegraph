package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Commit make request get update telegram
//func (call *GetUpdateCall) Commit() (*http.Response, []byte, error) {
//	prepareRequest := PrepareRequest{
//		Client:  call.Client,
//		Request: call.Request,
//	}
//	return prepareRequest.Commit()
//}




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
