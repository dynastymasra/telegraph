package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// GetMe prepare request get information about telegram bot
func (client *Client) GetMe() *PrepareRequest {
	url := BaseURL + fmt.Sprintf(EndpointGetMe, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		client:  client,
		request: request,
	}
}
