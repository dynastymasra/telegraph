package telegraph

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type (
	GetFileCall struct {
		Client  *Client
		Request *gorequest.SuperAgent
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
//func (call *GetFileCall) Download() (*http.Response, []byte, error) {
//	res, body, err := call.Commit()
//	if err != nil {
//		return nil, nil, err
//	}
//	result := &getFileResponse{}
//	if err := json.Unmarshal(body, result); err != nil {
//		return nil, nil, err
//	}
//	if res.StatusCode != http.StatusOK {
//		return nil, nil, fmt.Errorf(string(body))
//	}
//
//	return call.Client.GetContent(result.Result.FilePath).Commit()
//}

// Commit make request get file to telegram
//func (call *GetFileCall) Commit() (*http.Response, []byte, error) {
//	prepareRequest := PrepareRequest{
//		Client:  call.Client,
//		Request: call.Request,
//	}
//	return prepareRequest.Commit()
//}

// Download download random user profile photo
//func (call *GetUserProfilePhotoCall) Download() (*http.Response, []byte, error) {
//	res, body, err := call.Commit()
//	if err != nil {
//		return nil, nil, err
//	}
//	result := &getUserProfilePhotoResponse{}
//	if err := json.Unmarshal(body, result); err != nil {
//		return nil, nil, err
//	}
//	if res.StatusCode != http.StatusOK {
//		return nil, nil, fmt.Errorf(string(body))
//	}
//
//	var id string
//	var size int64
//	for _, first := range result.Result.Photos {
//		for _, second := range first {
//			if second.FileSize >= size {
//				id = second.FileID
//				size = second.FileSize
//			}
//		}
//	}
//
//	return call.Client.GetFile(id).Download()
//}

// GetContent used after call function GetFile, this download file from telegram
func (client *Client) GetContent(path string) *PrepareRequest {
	url := client.baseURL + fmt.Sprintf(EndpointGetContent, client.accessToken, path)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &PrepareRequest{
		Client:  client,
		Request: request,
	}
}
