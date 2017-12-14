package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// UserResponse struct to handle request and response telegram api
	UserResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetMe A simple method for testing your bot's auth token. Requires no parameters.
Returns basic information about the bot in form of a User object.
*/
func (client *Client) GetMe() *UserResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetMe, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &UserResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (user *UserResponse) Commit() (*User, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *User `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = user.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, user.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(user.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}
