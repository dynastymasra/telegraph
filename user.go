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

	// UserProfilePhotosResponse struct to handle request and response telegram api
	UserProfilePhotosResponse struct {
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
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *User `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, _, errs = user.Request.EndStruct(&model)
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

/*
GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
+ userId - Unique identifier of the target user

Available method can used with this method
+ SetOffset()
+ SetLimit()
*/
func (client *Client) GetUserProfilePhotos(userId int) *UserProfilePhotosResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetUserProfilePhoto, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).Query(fmt.Sprintf("user_id=%v", userId))

	return &UserProfilePhotosResponse{
		Client:  client,
		Request: request,
	}
}

// SetOffset Sequential number of the first photo to be returned. By default, all photos are returned.
func (user *UserProfilePhotosResponse) SetOffset(offset int) *UserProfilePhotosResponse {
	user.Request = user.Request.Query(fmt.Sprintf("offset=%v", offset))
	return user
}

// SetLimit Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
func (user *UserProfilePhotosResponse) SetLimit(limit int) *UserProfilePhotosResponse {
	user.Request = user.Request.Query(fmt.Sprintf("limit=%v", limit))
	return user
}

// Commit execute request to telegram
func (user *UserProfilePhotosResponse) Commit() (*UserProfilePhotos, *http.Response, error) {
	var errs []error
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *UserProfilePhotos `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, _, errs = user.Request.EndStruct(&model)
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
