package telegraph

import (
	"fmt"

	"net/http"

	"net/url"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// MessageResponse struct to handle request and response telegram api
	MessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// ArrayMessageResponse struct to handle request and array response telegram api
	ArrayMessageResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
SendMessage Use this method to send text messages. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ text - Text of the message to be sent

Available method can used with this method
+ SetParseMode()
+ SetDisableWebPagePreview()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendMessage(chatId interface{}, text string) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"text":    text,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendMessage, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetParseMode Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
func (message *MessageResponse) SetParseMode(mode string) *MessageResponse {
	body := JSON{
		"parse_mode": mode,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetDisableWebPagePreview Disables link previews for links in this message
func (message *MessageResponse) SetDisableWebPagePreview(disable bool) *MessageResponse {
	body := JSON{
		"disable_web_page_preview": disable,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
ForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ fromChatId - Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
+ messageId - Message identifier in the chat specified in from_chat_id

Available method can used with this method
+ SetDisableNotification()
*/
func (client *Client) ForwardMessage(chatId, fromChatId interface{}, messageId int) *MessageResponse {
	body := JSON{
		"chat_id":      chatId,
		"from_chat_id": fromChatId,
		"message_id":   messageId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointForwardMessage, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
SendPhoto Use this method to send photos. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ photo - Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
  pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data

Available method can used with this method
+ SetCaption()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendPhoto(chatId interface{}, photo string) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"photo":   photo,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendPhoto, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(photo); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(photo, "", "photo")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player.
Your audio must be in the .mp3 format. On success, the sent Message is returned.
Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
For sending voice messages, use the sendVoice method instead.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ audio - Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended),
  pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data.

Available method can used with this method
+ SetCaption()
+ SetDuration()
+ SetPerformer()
+ SetTitle()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendAudio(chatId interface{}, audio string) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"audio":   audio,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendAudio, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(audio); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(audio, "", "audio")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetPerformer Performer
func (message *MessageResponse) SetPerformer(performer string) *MessageResponse {
	body := JSON{
		"performer": performer,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetTitle Track name
func (message *MessageResponse) SetTitle(title string) *MessageResponse {
	body := JSON{
		"title": title,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
SendDocument Use this method to send general files. On success, the sent Message is returned.
Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ document - File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
  pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.

Available method can used with this method
+ SetCaption()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendDocument(chatId interface{}, document string) *MessageResponse {
	body := JSON{
		"chat_id":  chatId,
		"document": document,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendDocument, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(document); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(document, "", "document")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

/*
SendVideo Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
On success, the sent Message is returned.
Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ video - Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended),
  pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data.

Available method can used with this method
+ SetDuration()
+ SetWidth()
+ SetHeight()
+ SetCaption()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendVideo(chatId interface{}, video string) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"video":   video,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendVideo, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(video); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(video, "", "video")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetWidth Video width
func (message *MessageResponse) SetWidth(width int) *MessageResponse {
	body := JSON{
		"width": width,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetHeight Video height
func (message *MessageResponse) SetHeight(height int) *MessageResponse {
	body := JSON{
		"height": height,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
On success, the sent Message is returned.
Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ voice - Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
  pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.

Available method can used with this method
+ SetCaption()
+ SetDuration()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendVoice(chatId interface{}, voice string) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"voice":   voice,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendVoice, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(voice); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(voice, "", "voice")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetCaption Photo caption (may also be used when resending photos by file_id), 0-200 characters
func (message *MessageResponse) SetCaption(caption string) *MessageResponse {
	body := JSON{
		"caption": caption,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
SendVideoNote As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
Use this method to send video messages. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ videoNote - Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data.
  Sending video notes by a URL is currently unsupported

Available method can used with this method
+ SetDuration()
+ SetLength()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendVideoNote(chatId interface{}, videoNote string) *MessageResponse {
	body := JSON{
		"chat_id":    chatId,
		"video_note": videoNote,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendVideoNote, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(videoNote); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(videoNote, "", "video_note")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetLength Video width and height
func (message *MessageResponse) SetLength(length int) *MessageResponse {
	body := JSON{
		"length": length,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetDuration Duration of the audio in seconds
func (message *MessageResponse) SetDuration(duration int) *MessageResponse {
	body := JSON{
		"duration": duration,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
SendLocation Use this method to send point on the map. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ latitude - Latitude of the location
+ longitude - Longitude of the location

Available method can used with this method
+ SetLivePeriod()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendLocation(chatId interface{}, latitude, longitude float64) *MessageResponse {
	body := JSON{
		"chat_id":   chatId,
		"latitude":  latitude,
		"longitude": longitude,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetLivePeriod Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
func (message *MessageResponse) SetLivePeriod(livePeriod int) *MessageResponse {
	body := JSON{
		"livePeriod": livePeriod,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
SendVenue Use this method to send information about a venue. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ latitude - Latitude of the venue
+ longitude - Longitude of the venue
+ title - Name of the venue
+ address - Address of the venue

Available method can used with this method
+ SetFoursquareID()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendVenue(chatId interface{}, latitude, longitude float64, title, address string) *MessageResponse {
	body := JSON{
		"chat_id":   chatId,
		"latitude":  latitude,
		"longitude": longitude,
		"title":     title,
		"address":   address,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendVenue, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetFoursquareID Foursquare identifier of the venue
func (message *MessageResponse) SetFoursquareID(id string) *MessageResponse {
	body := JSON{
		"foursquare_id": id,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
SendContact Use this method to send phone contacts. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ phoneNumber - Contact's phone number
+ firstName - Contact's first name

Available method can used with this method
+ SetLastName()
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendContact(chatId interface{}, phoneNumber, firstName string) *MessageResponse {
	body := JSON{
		"chat_id":      chatId,
		"phone_number": phoneNumber,
		"first_name":   firstName,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendContact, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetLastName Contact's last name
func (message *MessageResponse) SetLastName(lastName string) *MessageResponse {
	body := JSON{
		"last_name": lastName,
	}
	message.Request = message.Request.Send(body)

	return message
}

/*
SendSticker Use this method to send .webp stickers. On success, the sent Message is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ sticker - Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
  pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data.

Available method can used with this method
+ SetDisableNotification()
+ SetReplyToMessageID()
+ SetForceReply()
+ SetInlineKeyboardMarkup()
+ SetReplyKeyboardMarkup()
+ SetReplyKeyboardRemove()
*/
func (client *Client) SendSticker(chatId interface{}, sticker string) *MessageResponse {
	body := JSON{
		"chat_id": chatId,
		"sticker": sticker,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendSticker, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(sticker); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(sticker, "", "video")
	}

	return &MessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (message *MessageResponse) SetDisableNotification(disable bool) *MessageResponse {
	body := JSON{
		"disable_notification": disable,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetReplyToMessageID If the message is a reply, ID of the original message
func (message *MessageResponse) SetReplyToMessageID(id int64) *MessageResponse {
	body := JSON{
		"reply_to_message_id": id,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetForceReply Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetForceReply(reply ForceReply) *MessageResponse {
	body := JSON{
		"reply_markup": reply,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetInlineKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *MessageResponse {
	body := JSON{
		"reply_markup": JSON{
			"inline_keyboard": inline,
		},
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetReplyKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetReplyKeyboardMarkup(reply ReplyKeyboardMarkup) *MessageResponse {
	body := JSON{
		"reply_markup": reply,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetReplyKeyboardRemove Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (message *MessageResponse) SetReplyKeyboardRemove(remove ReplyKeyboardRemove) *MessageResponse {
	body := JSON{
		"reply_markup": remove,
	}
	message.Request = message.Request.Send(body)

	return message
}

// Commit execute request to telegram
func (message *MessageResponse) Commit() (*Message, *http.Response, error) {
	var errs []error
	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *Message `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, _, errs = message.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, message.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(message.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}

/*
SendMediaGroup Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ media - A JSON-serialized array describing photos and videos to be sent, must include 2–10 items

Available method can used with this method
+ SetDisableNotification()
+ SetReplyToMessageID()
*/
func (client *Client) SendMediaGroup(chatId interface{}, media []InputMedia) *ArrayMessageResponse {
	body := JSON{
		"chat_id": chatId,
		"media":   media,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendMediaGroup, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &ArrayMessageResponse{
		Client:  client,
		Request: request,
	}
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (message *ArrayMessageResponse) SetDisableNotification(disable bool) *ArrayMessageResponse {
	body := JSON{
		"disable_notification": disable,
	}
	message.Request = message.Request.Send(body)

	return message
}

// SetReplyToMessageID If the message is a reply, ID of the original message
func (message *ArrayMessageResponse) SetReplyToMessageID(id int64) *ArrayMessageResponse {
	body := JSON{
		"reply_to_message_id": id,
	}
	message.Request = message.Request.Send(body)

	return message
}

// Commit execute request to telegram
func (message *ArrayMessageResponse) Commit() ([]Message, *http.Response, error) {
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result []Message `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, _, errs = message.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, message.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(message.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}
