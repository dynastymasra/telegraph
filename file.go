package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// FileResponse struct to handle request and response telegram api
	FileResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetFile Use this method to get basic info about a file and prepare it for downloading.
For the moment, bots can download files of up to 20MB in size. On success, a File object is returned.
The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>,
where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour.
When the link expires, a new one can be requested by calling getFile again.
This function only return json value not file, download file use function GetContent()
+ fileId - File identifier to get info about
*/
func (client *Client) GetFile(fileId string) *FileResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetFile, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).Query(fmt.Sprintf("file_id=%v", fileId))

	return &FileResponse{
		Client:  client,
		Request: request,
	}
}

/*
UploadStickerFile Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times).
Returns the uploaded File on success.
+ userId - User identifier of sticker file owner
+ pngSticker - Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px.
*/
func (client *Client) UploadStickerFile(userId int64, pngSticker string) *FileResponse {
	body := JSON{
		"user_id": userId,
	}
	url := client.baseURL + fmt.Sprintf(EndpointUploadStickerFile, client.accessToken)
	request := gorequest.New().Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body).Type(gorequest.TypeMultipart).
		SendFile(pngSticker, "", "png_sticker")

	return &FileResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (user *FileResponse) Commit() (*File, *http.Response, error) {
	var errs []error
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *File `json:"result,omitempty"`
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
