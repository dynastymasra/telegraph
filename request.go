package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type (
	JSON map[string]interface{}
)

func (client *Client) SendChatAction(chatID, action string) *VoidResponse {
	body := JSON{
		"chat_id": chatID,
		"action":  action,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendChatAction, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}
