package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// ChatResponse struct to handle request and response telegram api
	ChatResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetChat Use this method to get up to date information about the chat (current name of the user for one-on-one conversations,
current username of a user, group or channel, etc.). Returns a Chat object on success.
*/
func (client *Client) GetChat(chatId interface{}) *ChatResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetChat, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatId))

	return &ChatResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (void *ChatResponse) Commit() (*Chat, *http.Response, error) {
	var body []byte
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *Chat `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = void.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, void.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(void.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}
