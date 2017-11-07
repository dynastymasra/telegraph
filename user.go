package telegraph

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	UserResponse struct {
		Client   *Client
		Request  *gorequest.SuperAgent
		err      error
		Response *http.Response
		Body     []byte
	}
)

// GetMe A simple method for testing your bot's auth token. Requires no parameters.
// Returns basic information about the bot in form of a User object.
func (client *Client) GetMe() *UserResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetMe, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &UserResponse{
		Client:  client,
		Request: request,
	}
}

// Commit request to telegram api
func (user *UserResponse) Commit() *UserResponse {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = user.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, user.Client.expBackOff); err != nil {
		user.Response = nil
		user.Body = nil
		user.err = err
		return user
	}

	user.Response = res
	user.Body = body
	user.err = nil
	return user
}

// Parse response get me to struct
func (user *UserResponse) Parse() (*User, int, error) {
	if user.err != nil {
		return nil, http.StatusInternalServerError, user.err
	}

	model := struct {
		ErrorResponse
		Result *User `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(user.Body, &model); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if user.Response.StatusCode != http.StatusOK {
		return nil, model.ErrorCode, fmt.Errorf(model.Description)
	}
	return model.Result, user.Response.StatusCode, nil
}
