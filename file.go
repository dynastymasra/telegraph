package telegraph

import (
	"fmt"

	"encoding/json"
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

	operation := func() error {
		res, body, errs = file.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, file.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(file.Request), err
	}

	return parseFile(res, body)
}

func parseFile(res *http.Response, body []byte) (*File, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *File `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}

	return model.Result, res, nil
}

// Download direct download file from telegram after get file
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
