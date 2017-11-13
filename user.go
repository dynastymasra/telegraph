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

	UserProfilePhotos struct {
		TotalCount int           `json:"total_count"`
		Photos     [][]PhotoSize `json:"photos"`
	}

	UserResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	UserProfilePhotosResponse struct {
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
		return nil, MakeHTTPResponse(user.Request), err
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

/*
GetUserProfilePhotos Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
*/
func (client *Client) GetUserProfilePhotos(userID int64) *UserProfilePhotosResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetUserProfilePhoto, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).Query(fmt.Sprintf("user_id=%v", userID))

	return &UserProfilePhotosResponse{
		Client:  client,
		Request: request,
	}
}

// SetLimit Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
func (user *UserProfilePhotosResponse) SetLimit(limit int) *UserProfilePhotosResponse {
	return &UserProfilePhotosResponse{
		Client:  user.Client,
		Request: user.Request.Query(fmt.Sprintf("limit=%v", limit)),
	}
}

// SetOffset Sequential number of the first photo to be returned. By default, all photos are returned.
func (user *UserProfilePhotosResponse) SetOffset(offset int) *UserProfilePhotosResponse {
	return &UserProfilePhotosResponse{
		Client:  user.Client,
		Request: user.Request.Query(fmt.Sprintf("offset=%v", offset)),
	}
}

// Commit request to telegram api
func (user *UserProfilePhotosResponse) Commit() (*UserProfilePhotos, *http.Response, error) {
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
		return nil, MakeHTTPResponse(user.Request), err
	}

	return parseUserProfilePhotos(res, body)
}

func parseUserProfilePhotos(res *http.Response, body []byte) (*UserProfilePhotos, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *UserProfilePhotos `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}

	return model.Result, res, nil
}

// Download direct get file and download file from telegram
func (user *UserProfilePhotosResponse) Download() (*http.Response, []byte, error) {
	profile, _, err := user.Commit()
	if err != nil {
		return nil, nil, err
	}

	var fileID string
	var fileSize int
	for _, first := range profile.Photos {
		for _, second := range first {
			if second.FileSize > fileSize {
				fileID = second.FileID
				fileSize = second.FileSize
			}
		}
		if len(fileID) > 0 {
			break
		}
	}

	return user.Client.GetFile(fileID).Download()
}
