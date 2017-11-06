package telegraph

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	Info struct {
		Client   *Client
		Request  *gorequest.SuperAgent
		err      error
		Response *http.Response
		Body     []byte
	}
)

// GetMe A simple method for testing your bot's auth token. Requires no parameters.
// Returns basic information about the bot in form of a User object.
func (client *Client) GetMe() *Info {
	url := client.baseURL + fmt.Sprintf(EndpointGetMe, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &Info{
		Client:  client,
		Request: request,
	}
}

// Commit request to telegram api
func (info *Info) Commit() *Info {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = info.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, info.Client.expBackOff); err != nil {
		info.Response = nil
		info.Body = nil
		info.err = err
		return info
	}

	info.Response = res
	info.Body = body
	info.err = nil
	return info
}

// Parse response get me to struct
func (info *Info) Parse() (*User, int, error) {
	if info.err != nil {
		return nil, http.StatusInternalServerError, info.err
	}

	model := struct {
		ErrorResponse
		Result *User `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(info.Body, &model); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if info.Response.StatusCode != http.StatusOK {
		return nil, model.ErrorCode, fmt.Errorf(model.Description)
	}
	return model.Result, info.Response.StatusCode, nil
}
