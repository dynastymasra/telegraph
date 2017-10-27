package telegraph

import (
	"fmt"

	"net/http"

	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

type (
	GetFileCall struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	GetUserProfilePhotoCall struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	getUserProfilePhotoResponse struct {
		Result      photos `json:"photos,omitempty"`
		Description string `json:"description,omitempty"`
	}

	photos struct {
		Photos [][]result `json:"photos"`
	}

	getFileResponse struct {
		OK          bool   `json:"ok"`
		Result      result `json:"result,omitempty"`
		Description string `json:"description,omitempty"`
	}

	result struct {
		FileID   string `json:"file_id"`
		FileSize int64  `json:"file_size"`
		FilePath string `json:"file_path"`
	}
)

// GetFile get path file from telegram
func (client *Client) GetFile(fileId string) *GetFileCall {
	url := client.baseURL + fmt.Sprintf(EndpointGetFile, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("file_id=%v", fileId))

	return &GetFileCall{
		Client:  client,
		Request: request,
	}
}

// Download used for direct download file after call GetFile
func (call *GetFileCall) Download() (*http.Response, []byte, error) {
	res, body, err := call.Commit()
	if err != nil {
		return nil, nil, err
	}
	result := &getFileResponse{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf(string(body))
	}

	return call.Client.GetContent(result.Result.FilePath).Commit()
}

// Commit make request get file to telegram
func (call *GetFileCall) Commit() (*http.Response, []byte, error) {
	prepareRequest := PrepareRequest{
		Client:  call.Client,
		Request: call.Request,
	}
	return prepareRequest.Commit()
}

// GetUserProfilePhoto prepare get user profile photo
func (client *Client) GetUserProfilePhoto(userID string) *GetUserProfilePhotoCall {
	url := client.baseURL + fmt.Sprintf(EndpointGetUserProfilePhoto, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("user_id=%v", userID))

	return &GetUserProfilePhotoCall{
		Client:  client,
		Request: request,
	}
}

// Limit set limit response returned
func (call *GetUserProfilePhotoCall) Limit(limit int) *GetUserProfilePhotoCall {
	return &GetUserProfilePhotoCall{
		Client:  call.Client,
		Request: call.Request.Query(fmt.Sprintf("limit=%v", limit)),
	}
}

// Offset set offset response returned
func (call *GetUserProfilePhotoCall) Offset(offset int) *GetUserProfilePhotoCall {
	return &GetUserProfilePhotoCall{
		Client:  call.Client,
		Request: call.Request.Query(fmt.Sprintf("offset=%v", offset)),
	}
}

// Download download random user profile photo
func (call *GetUserProfilePhotoCall) Download() *PrepareRequest {
	res, body, err := call.Commit()
	if err != nil {
		return nil
	}
	result := &getUserProfilePhotoResponse{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil
	}
	if res.StatusCode != http.StatusOK {
		return nil
	}

	var path string
	for _, first := range result.Result.Photos {
		for _, second := range first {
			if len(second.FilePath) > 0 {
				path = second.FilePath
				break
			}
			if len(path) > 0 {
				break
			}
		}
		if len(path) > 0 {
			break
		}
	}

	return call.Client.GetContent(path)
}

// Commit get user profile photo
func (call *GetUserProfilePhotoCall) Commit() (*http.Response, []byte, error) {
	prepareRequest := PrepareRequest{
		Client:  call.Client,
		Request: call.Request,
	}
	return prepareRequest.Commit()
}

// GetContent used after call function GetFile, this download file from telegram
func (client *Client) GetContent(path string) *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(EndpointGetContent, client.accessToken, path)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		Client:  client,
		Request: request,
	}
}
