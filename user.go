package telegraph

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// User This object represents a Telegram user or bot.
	User struct {
		ID           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name,omitempty"`
		Username     string `json:"username,omitempty"`
		LanguageCode string `json:"language_code,omitempty"`
	}

	UserResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
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
func (user *UserResponse) Commit() (*User, *http.Response, error) {
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
		return nil, &http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return parseUser(res, body)
}

func parseUser(res *http.Response, body []byte) (*User, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *User `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}

	return model.Result, res, nil
}
