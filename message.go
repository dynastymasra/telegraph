package telegraph

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup               = "group"
	ChatTypeSuperGroup          = "supergroup"
	ChatTypeChannel             = "channel"
)

type (
	Message struct {
		MessageID int64       `json:"message_id"`
		Date      int64       `json:"date"`
		Chat      Chat        `json:"chat"`
		From      *From       `json:"from,omitempty"`
		Text      string      `json:"text,omitempty"`
		Photos    []PhotoSize `json:"photo,omitempty"`
		Sticker   *Sticker    `json:"sticker,omitempty"`
	}

	MessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
		url     string
	}

	PhotoSize struct {
		FileID   string `json:"file_id"`
		Width    int64  `json:"width"`
		Height   int64  `json:"height"`
		FileSize int64  `json:"file_size"`
	}

	Sticker struct {
		FileID       string        `json:"file_id"`
		Width        int64         `json:"width"`
		Height       int64         `json:"height"`
		Thumb        *PhotoSize    `json:"thumb,omitempty"`
		Emoji        string        `json:"emoji,omitempty"`
		SetName      string        `json:"set_name,omitempty"`
		FileSize     int64         `json:"file_size,omitempty"`
		MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	}

	MaskPosition struct {
		Point  string  `json:"point"`
		XShift float64 `json:"x_shift"`
		YShift float64 `json:"y_shift"`
		Scale  float64 `json:"scale"`
	}

	Chat struct {
		ID                     int64    `json:"id"`
		Type                   ChatType `json:"type"`
		Title                  string   `json:"title,omitempty"`
		Username               string   `json:"username,omitempty"`
		FirstName              string   `json:"first_name,omitempty"`
		LastName               string   `json:"last_name,omitempty"`
		AllMemberAdministrator bool     `json:"all_members_are_administrators,omitempty"`
		Description            string   `json:"description,omitempty"`
		InviteLink             string   `json:"invite_link,omitempty"`
	}

	From struct {
		ID           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name,omitempty"`
		Username     string `json:"username,omitempty"`
		LanguageCode string `json:"language_code,omitempty"`
	}
)

// SendMessage Use this method to send text messages. On success, the sent Message is returned.
func (client *Client) SendMessage(message SendMessage) *MessageResponse {
	url := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// ForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
func (client *Client) ForwardMessage(message ForwardMessage) *MessageResponse {
	url := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (client *Client) SendPhoto(message SendPhoto, upload bool) *MessageResponse {
	endpoint := client.baseURL + fmt.Sprintf(message.endpoint, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(message)

	if upload {
		request.Type(gorequest.TypeMultipart).SendFile(message.Photo, "", "photo")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
		url:     message.Photo,
	}
}

// Commit process request send message to telegram
func (message *MessageResponse) Commit() (*Message, *http.Response, error) {
	var errs []error
	var body []byte
	res := &http.Response{}

	operation := func() error {
		res, body, errs = message.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, message.Client.expBackOff); err != nil {
		return nil, makeHTTPResponse(message.Request), err
	}
	return parseMessage(res, body)
}

func parseMessage(res *http.Response, body []byte) (*Message, *http.Response, error) {
	model := struct {
		ErrorResponse
		Result *Message `json:"result,omitempty"`
	}{}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, res, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf(model.Description)
	}
	return model.Result, res, nil
}
