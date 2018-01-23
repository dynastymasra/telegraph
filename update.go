package telegraph

import (
	"encoding/json"

	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// ArrayUpdateResponse struct to handle request and response from telegram api with array update
	ArrayUpdateResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

// WebHookParseRequest function for parse request from telegram web hook, return struct Update if success
func WebHookParseRequest(r []byte) (*Update, error) {
	update := &Update{}

	if err := json.Unmarshal(r, update); err != nil {
		return nil, err
	}

	return update, nil
}

/*
GetUpdates Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.

Notes
1. This method will not work if an outgoing webhook is set up.
2. In order to avoid getting duplicate updates, recalculate offset after each server response
*/
func (client *Client) GetUpdates() *ArrayUpdateResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetUpdate, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &ArrayUpdateResponse{
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
func (update *ArrayUpdateResponse) SetOffset(offset int) *ArrayUpdateResponse {
	update.Request = update.Request.Query(fmt.Sprintf("offset=%v", offset))
	return update
}

// SetLimit Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
func (update *ArrayUpdateResponse) SetLimit(limit int) *ArrayUpdateResponse {
	update.Request = update.Request.Query(fmt.Sprintf("limit=%v", limit))
	return update
}

// SetTimeout Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling.
// Should be positive, short polling should be used for testing purposes only.
func (update *ArrayUpdateResponse) SetTimeout(timeout int) *ArrayUpdateResponse {
	update.Request = update.Request.Query(fmt.Sprintf("timeout=%v", timeout))
	return update
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
func (update *ArrayUpdateResponse) SetAllowedUpdates(updates ...string) *ArrayUpdateResponse {
	update.Request = update.Request.Query(fmt.Sprintf("allowed_updates=%v", updates))
	return update
}

// Commit request to telegram api
func (update *ArrayUpdateResponse) Commit() ([]Update, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result []Update `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = update.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, update.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(update.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}
