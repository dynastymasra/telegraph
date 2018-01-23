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

	// ChatMemberResponse struct to handle request and response telegram api
	ChatMemberResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// ArrayChatMemberResponse struct to handle request and response telegram api
	ArrayChatMemberResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetChat Use this method to get up to date information about the chat (current name of the user for one-on-one conversations,
current username of a user, group or channel, etc.). Returns a Chat object on success.
+ chatId - Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
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

/*
GetChatAdministrator Use this method to get a list of administrators in a chat.
On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots.
If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
+ chatId - Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
*/
func (client *Client) GetChatAdministrator(chatId interface{}) *ArrayChatMemberResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetChatAdministrators, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatId))

	return &ArrayChatMemberResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (void *ArrayChatMemberResponse) Commit() ([]ChatMember, *http.Response, error) {
	var body []byte
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result []ChatMember `json:"result,omitempty"`
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

/*
GetChatMember Use this method to get information about a member of a chat. Returns a ChatMember object on success.
+ chatId - Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
+ userId - Unique identifier of the target user
*/
func (client *Client) GetChatMember(chatId interface{}, userId int64) *ChatMemberResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v&user_id=%v", chatId, userId))

	return &ChatMemberResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (void *ChatMemberResponse) Commit() (*ChatMember, *http.Response, error) {
	var body []byte
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *ChatMember `json:"result,omitempty"`
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
