package telegraph

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	ChatType   string
	StatusType string
)

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup               = "group"
	ChatTypeSuperGroup          = "supergroup"
	ChatTypeChannel             = "channel"

	StatusTypeCreator       StatusType = "creator"
	StatusTypeAdministrator            = "administrator"
	StatusTypeMember                   = "member"
	StatusTypeRestricted               = "restricted"
	StatusTypeLeft                     = "left"
	StatusTypeKicked                   = "kicked"
)

type (
	// Chat This object represents a chat.
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

	// ChatPhoto This object represents a chat photo.
	ChatPhoto struct {
		SmallFileID string `json:"small_file_id"`
		BigFileID   string `json:"big_file_id"`
	}

	// ChatResponse struct to handle request and response telegram api
	ChatResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// ChatMember This object contains information about one member of a chat.
	ChatMember struct {
		User                 *User      `json:"user"`
		Status               StatusType `json:"status"`
		UntilDate            int64      `json:"until_date,omitempty"`
		CanBeEdited          bool       `json:"can_be_edited,omitempty"`
		CanChangeInfo        bool       `json:"can_change_info,omitempty"`
		CanPostMessage       bool       `json:"can_post_messages,omitempty"`
		CanEditMessage       bool       `json:"can_edit_messages,omitempty"`
		CanDeleteMessage     bool       `json:"can_delete_messages,omitempty"`
		CanInviteUser        bool       `json:"can_invite_users,omitempty"`
		CanRestrictMember    bool       `json:"can_restrict_members,omitempty"`
		CanPinMessage        bool       `json:"can_pin_messages,omitempty"`
		CanPromoteMember     bool       `json:"can_promote_members,omitempty"`
		CanSendMessage       bool       `json:"can_send_messages,omitempty"`
		CanSendMediaMessage  bool       `json:"can_send_media_messages,omitempty"`
		CanSendOtherMessage  bool       `json:"can_send_other_messages,omitempty"`
		CanAddWebPagePreview bool       `json:"can_add_web_page_previews,omitempty"`
	}

	// ChatMemberResponse struct to handle request and response telegram api
	ChatMemberResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// ChatMemberArrayResponse struct to handle request and response telegram api
	ChatMemberArrayResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// IntegerResponse struct to handle request and response telegram api
	IntegerResponse struct {
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

/*
GetChatAdministrators Use this method to get a list of administrators in a chat.
On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots.
If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
*/
func (client *Client) GetChatAdministrators(chatID interface{}) *ChatMemberArrayResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetChatAdministrators, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatID))

	return &ChatMemberArrayResponse{
		Client:  client,
		Request: request,
	}
}

// Commit request to telegram api
func (chat *ChatMemberArrayResponse) Commit() ([]ChatMember, *http.Response, error) {
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

	return parseArrayChatMember(res, body)
}

func parseArrayChatMember(res *http.Response, body []byte) ([]ChatMember, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result []ChatMember `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}

	return model.Result, res, nil
}

/*
GetChatMembersCount Use this method to get the number of members in a chat. Returns Int on success.
*/
func (client *Client) GetChatMembersCount(chatID interface{}) *IntegerResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetChatMembersCount, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatID))

	return &IntegerResponse{
		Client:  client,
		Request: request,
	}
}

// Commit request to telegram api
func (chat *IntegerResponse) Commit() (*int, *http.Response, error) {
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

	return parseIntegerResponse(res, body)
}

func parseIntegerResponse(res *http.Response, body []byte) (*int, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *int `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}

	return model.Result, res, nil
}

/*
GetChatMember Use this method to get information about a member of a chat. Returns a ChatMember object on success.
*/
func (client *Client) GetChatMember(chatID interface{}, userID int64) *ChatMemberResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetChatMember, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v&user_id=%v", chatID, userID))

	return &ChatMemberResponse{
		Client:  client,
		Request: request,
	}
}

// Commit request to telegram api
func (chat *ChatMemberResponse) Commit() (*ChatMember, *http.Response, error) {
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

	return parseChatMemberResponse(res, body)
}

func parseChatMemberResponse(res *http.Response, body []byte) (*ChatMember, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *ChatMember `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}

	return model.Result, res, nil
}
