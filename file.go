package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// File This object represents a file ready to be downloaded.
	// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
	// It is guaranteed that the link will be valid for at least 1 hour.
	// When the link expires, a new one can be requested by calling getFile.
	File struct {
		FileID   string `json:"file_id"`
		FileSize int64  `json:"file_size,omitempty"`
		FilePath string `json:"file_path,omitempty"`
	}

	// FileResponse struct to handle request and response telegram api
	FileResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetFile This object represents a file ready to be downloaded.
The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
It is guaranteed that the link will be valid for at least 1 hour.
When the link expires, a new one can be requested by calling getFile.
*/
func (client *Client) GetFile(fileId string) *FileResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetFile, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("file_id=%v", fileId))

	return &FileResponse{
		Client:  client,
		Request: request,
	}
}

// Commit request to telegram api
func (file *FileResponse) Commit() (*File, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *File `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = file.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, file.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(file.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}

// Download direct download file from telegram after get file, cannot used with Commit
func (file *FileResponse) Download() (*http.Response, []byte, error) {
	model, _, err := file.Commit()
	if err != nil {
		return nil, nil, err
	}

	return file.Client.GetContent(model.FilePath).Download()
}

// GetContent used after call function GetFile, this download file from telegram with path file
func (client *Client) GetContent(path string) *VoidResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetContent, client.accessToken, path)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}
