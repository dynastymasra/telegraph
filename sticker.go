package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// StickerSetResponse struct to handle request and response telegram api
	StickerSetResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetStickerSet Use this method to get a sticker set. On success, a StickerSet object is returned.
+ name - Use this method to get a sticker set. On success, a StickerSet object is returned.
*/
func (client *Client) GetStickerSet(name string) *StickerSetResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetStickerSet, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("name=%v", name))

	return &StickerSetResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (sticker *StickerSetResponse) Commit() (*StickerSet, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *StickerSet `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = sticker.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, sticker.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(sticker.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}
