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
		Err     error
	}

	GetUserProfilePhotoCall struct {
		Client  *Client
		Request *gorequest.SuperAgent
		Err     error
	}

	getUserProfilePhotoResponse struct {
		Result photos `json:"photos,omitempty"`
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
		Err:     nil,
	}
}

// Download used for direct download file after call GetFile
func (call *GetFileCall) Download() *GetFileCall {
	res, body, err := call.Commit()
	if err != nil {
		return &GetFileCall{
			Err: err,
		}
	}
	result := &getFileResponse{}
	if err := json.Unmarshal(body, result); err != nil {
		return &GetFileCall{
			Err: err,
		}
	}
	if res.StatusCode != http.StatusOK {
		return &GetFileCall{
			Err: fmt.Errorf(result.Description),
		}
	}

	prepareRequest := call.Client.GetContent(result.Result.FilePath)
	return &GetFileCall{
		Client:  prepareRequest.client,
		Request: prepareRequest.request,
		Err:     nil,
	}
}

// Commit make request get file to telegram
func (call *GetFileCall) Commit() (*http.Response, []byte, error) {
	if call.Err != nil {
		return nil, nil, call.Err
	}

	prepareRequest := PrepareRequest{
		client:  call.Client,
		request: call.Request,
	}
	return prepareRequest.Commit()
}

// GetUserProfilePhoto prepare get user profile photo
func (client *Client) GetUserProfilePhoto(userID string) *GetUserProfilePhotoCall {
	url := client.baseURL + fmt.Sprintf(EndpointGetUserProfilePhoto, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("user_id=%v", userID))

	return &GetUserProfilePhotoCall{
		Client:client,
		Request:request,
		Err:nil,
	}
}

// Download download random user profile photo
func (call *GetUserProfilePhotoCall) Download() *GetUserProfilePhotoCall {
	res, body, err := call.Commit()
	if err != nil {
		return &GetUserProfilePhotoCall{
			Err: err,
		}
	}
	result := &getUserProfilePhotoResponse{}
	if err := json.Unmarshal(body, result); err != nil {
		return &GetUserProfilePhotoCall{
			Err: err,
		}
	}
	if res.StatusCode != http.StatusOK {
		return &GetUserProfilePhotoCall{
			Err: fmt.Errorf(result.Description),
		}
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
	prepareRequest := call.Client.GetContent(path)
	return &GetUserProfilePhotoCall{
		Client:  prepareRequest.client,
		Request: prepareRequest.request,
		Err:     nil,
	}
}

// Commit get user profile photo
func (call *GetUserProfilePhotoCall) Commit() (*http.Response, []byte, error) {
	if call.Err != nil {
		return nil, nil, call.Err
	}

	prepareRequest := PrepareRequest{
		client:  call.Client,
		request: call.Request,
	}
	return prepareRequest.Commit()
}

// GetContent used after call function GetFile, this download file from telegram
func (client *Client) GetContent(path string) *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(EndpointGetContent, client.accessToken, path)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		client:  client,
		request: request,
	}
}
