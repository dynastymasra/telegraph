package telegraph

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup               = "group"
	ChatTypeSuperGroup          = "supergroup"
	ChatTypeChannel             = "channel"
)

type (
	Chat struct {
		ID                     int64      `json:"id"`
		Type                   ChatType   `json:"type"`
		Title                  string     `json:"title,omitempty"`
		Username               string     `json:"username,omitempty"`
		FirstName              string     `json:"first_name,omitempty"`
		LastName               string     `json:"last_name,omitempty"`
		AllMemberAdministrator bool       `json:"all_members_are_administrators,omitempty"`
		Photo                  *ChatPhoto `json:"photo,omitempty"`
		Description            string     `json:"description,omitempty"`
		InviteLink             string     `json:"invite_link,omitempty"`
		PinnedMessage          *Message   `json:"pinned_message,omitempty"`
		StickerSetName         string     `json:"sticker_set_name,omitempty"`
		CanSetStickerSet       bool       `json:"can_set_sticker_set,omitempty"`
	}

	ChatPhoto struct {
		SmallFileID string `json:"small_file_id"`
		BigFileID   string `json:"big_file_id"`
	}

	ChatResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
GetChat Use this method to get up to date information about the chat (current name of the user for one-on-one conversations,
current username of a user, group or channel, etc.). Returns a Chat object on success.
*/
func (client *Client) GetChat(chatID interface{}) *ChatResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetChat, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatID))

	return &ChatResponse{
		Client:  client,
		Request: request,
	}
}

// Commit request to telegram api
func (chat *ChatResponse) Commit() (*Chat, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = chat.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, chat.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(chat.Request), err
	}

	return parseChat(res, body)
}

func parseChat(res *http.Response, body []byte) (*Chat, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *Chat `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}

	return model.Result, res, nil
}
