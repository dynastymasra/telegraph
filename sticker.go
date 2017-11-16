package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// StickerSet This object represents a sticker set.
	StickerSet struct {
		Name          string    `json:"name"`
		Title         string    `json:"title"`
		ContainsMasks bool      `json:"contains_masks"`
		Stickers      []Sticker `json:"stickers"`
	}

	// Sticker This object represents a sticker.
	Sticker struct {
		FileID       string        `json:"file_id"`
		Width        int           `json:"width"`
		Height       int           `json:"height"`
		Thumb        *PhotoSize    `json:"thumb,omitempty"`
		Emoji        string        `json:"emoji,omitempty"`
		SetName      string        `json:"set_name,omitempty"`
		FileSize     int           `json:"file_size,omitempty"`
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	}

	// StickerSetResponse struct to handle request and response telegram api
	StickerSetResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetStickerSet Use this method to get a sticker set. On success, a StickerSet object is returned.
*/
func (client *Client) GetStickerSet(name string) *StickerSetResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetStickerSet, client.accessToken)
	request := gorequest.New().Get(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("name=%v", name))

	return &StickerSetResponse{
		Client:  client,
		Request: request,
	}
}

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
