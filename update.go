package telegraph

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// Update This object represents an incoming update.
	// At most one of the optional parameters can be present in any given update.
	Update struct {
		UpdateID           int64               `json:"update_id"`
		Message            *Message            `json:"message,omitempty"`
		EditedMessage      *Message            `json:"edited_message,omitempty"`
		ChannelPost        *Message            `json:"channel_post,omitempty"`
		EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
		InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
		ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
		CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
		ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
		PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	}

	UpdateArrayResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

// GetUpdate Use this method to receive incoming updates using long polling (wiki).
// An Array of Update objects is returned.
// Note: This method will not work if an outgoing webhook is set up.
// In order to avoid getting duplicate updates, recalculate offset after each server response.
func (client *Client) GetUpdate() *UpdateArrayResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetUpdate, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &UpdateArrayResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetOffset Identifier of the first update to be returned.
Must be greater by one than the highest among the identifiers of previously received updates.
By default, updates starting with the earliest unconfirmed update are returned.
An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id.
The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue.
All previous updates will forgotten.
*/
func (update *UpdateArrayResponse) SetOffset(offset int) *UpdateArrayResponse {
	return &UpdateArrayResponse{
		Client:  update.Client,
		Request: update.Request.Query(fmt.Sprintf("offset=%v", offset)),
	}
}

// SetLimit Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
func (update *UpdateArrayResponse) SetLimit(limit int) *UpdateArrayResponse {
	return &UpdateArrayResponse{
		Client:  update.Client,
		Request: update.Request.Query(fmt.Sprintf("limit=%v", limit)),
	}
}

// SetTimeout Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling.
// Should be positive, short polling should be used for testing purposes only.
func (update *UpdateArrayResponse) SetTimeout(timeout int) *UpdateArrayResponse {
	return &UpdateArrayResponse{
		Client:  update.Client,
		Request: update.Request.Query(fmt.Sprintf("timeout=%v", timeout)),
	}
}

/*
SetAllowedUpdates List the types of updates you want your bot to receive.
For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types.
See Update for a complete list of available update types.
Specify an empty list to receive all updates regardless of type (default).
If not specified, the previous setting will be used.

Please note that this parameter doesn't affect updates created before the call to the getUpdates,
so unwanted updates may be received for a short period of time.
*/
func (update *UpdateArrayResponse) SetAllowedUpdates(updates ...string) *UpdateArrayResponse {
	return &UpdateArrayResponse{
		Client:  update.Client,
		Request: update.Request.Query(fmt.Sprintf("allowed_updates=%v", updates)),
	}
}

// Commit request to telegram api
func (update *UpdateArrayResponse) Commit() ([]Update, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = update.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, update.Client.expBackOff); err != nil {
		return nil, &http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return parseArrayUpdate(res, body)
}

func parseArrayUpdate(res *http.Response, body []byte) ([]Update, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result []Update `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}
	return model.Result, res, nil
}
