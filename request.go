package telegraph

import (
	"fmt"

	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// VoidResponse struct to handle request and response telegram api
	VoidResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}
)

/*
SetWebHook Use this method to specify a url and receive incoming updates via an outgoing webhook.
Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
containing a JSON-serialized Update. In case of an unsuccessful request,
we will give up after a reasonable amount of attempts. Returns true.
*/
func (client *Client) SetWebHook(webHook string) *VoidResponse {
	body := JSON{
		"url": webHook,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSetWebHook, client.accessToken)
	request := gorequest.New().Post(url).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetCertificate Upload your public key certificate so that the root certificate in use can be checked.
// See our self-signed guide for details.
func (void *VoidResponse) SetCertificate(path string) *VoidResponse {
	void.Request = void.Request.Type(gorequest.TypeMultipart).SendFile(path, "", "certificate")

	return void
}

// SetMaxConnection Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery,
// 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server,
// and higher values to increase your bot’s throughput.
func (void *VoidResponse) SetMaxConnection(conn int) *VoidResponse {
	body := JSON{
		"max_connections": conn,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
SetAllowedUpdates List the types of updates you want your bot to receive.
For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types.
See Update for a complete list of available update types.
Specify an empty list to receive all updates regardless of type (default).
If not specified, the previous setting will be used.
*/
func (void *VoidResponse) SetAllowedUpdates(allowed ...string) *VoidResponse {
	body := JSON{
		"allowed_updates": allowed,
	}
	void.Request = void.Request.Send(body)

	return void
}

// DeleteWebHook Use this method to remove webhook integration if you decide to switch back to getUpdates.
// Returns True on success. Requires no parameters.
func (client *Client) DeleteWebHook() *VoidResponse {
	url := client.baseURL + fmt.Sprintf(EndpointDeleteWebHook, client.accessToken)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
EditMessageLiveLocation Use this method to edit live location messages sent by the bot or via the bot (for inline bots).
A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
*/
func (client *Client) EditMessageLiveLocation(latitude, longitude float64) *VoidResponse {
	body := JSON{
		"latitude":  latitude,
		"longitude": longitude,
	}

	url := client.baseURL + fmt.Sprintf(EndpointEditMessageLiveLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
StopMessageLiveLocation Use this method to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires.
On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
*/
func (client *Client) StopMessageLiveLocation() *VoidResponse {
	url := client.baseURL + fmt.Sprintf(EndpointStopMessageLiveLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetChatID Required if inline_message_id is not specified.
// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (void *VoidResponse) SetChatID(chatId interface{}) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetMessageID Required if inline_message_id is not specified. Identifier of the sent message
func (void *VoidResponse) SetMessageID(messageId int) *VoidResponse {
	body := JSON{
		"message_id": messageId,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetInlineMessageID Required if chat_id and message_id are not specified. Identifier of the inline message
func (void *VoidResponse) SetInlineMessageID(inlineMessage string) *VoidResponse {
	body := JSON{
		"inline_message_id": inlineMessage,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetInlineKeyboardMarkup Additional interface options. A JSON-serialized object for an inline keyboard,
// custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
func (void *VoidResponse) SetInlineKeyboardMarkup(inline [][]InlineKeyboardButton) *VoidResponse {
	body := JSON{
		"reply_markup": JSON{
			"inline_keyboard": inline,
		},
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
SendChatAction Use this method when you need to tell the user that something is happening on the bot's side.
The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
Returns True on success.

Example: The ImageBot needs some time to process a request and upload the image.
Instead of sending a text message along the lines of “Retrieving image, please wait…”,
the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.

Type of action to broadcast.
Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos,
record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files,
find_location for location data, record_video_note or upload_video_note for video notes.
*/
func (client *Client) SendChatAction(chatId interface{}, action string) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"action":  action,
	}

	url := client.baseURL + fmt.Sprintf(EndpointSendChatAction, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
KickChatMember Use this method to kick a user from a group, a supergroup or a channel.
In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc.,
unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns True on success.

Note: In regular groups (non-supergroups), this method will only work if the ‘All Members Are Admins’ setting is off in the target group.
Otherwise members may only be removed by the group's creator or by the member that added them.
*/
func (client *Client) KickChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}

	url := client.baseURL + fmt.Sprintf(EndpointKickChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
RestrictChatMember Use this method to restrict a user in a supergroup.
The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.
*/
func (client *Client) RestrictChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}

	url := client.baseURL + fmt.Sprintf(EndpointRestrictChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetUntilDate Date when the user will be unbanned, unix time.
// If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
func (void *VoidResponse) SetUntilDate(date int64) *VoidResponse {
	body := JSON{
		"until_date": date,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanSendMessages Pass True, if the user can send text messages, contacts, locations and venues
func (void *VoidResponse) SetCanSendMessage(can bool) *VoidResponse {
	body := JSON{
		"can_send_messages": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanSendMediaMessages Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes,
// implies can_send_messages
func (void *VoidResponse) SetCanSendMediaMessage(can bool) *VoidResponse {
	body := JSON{
		"can_send_media_messages": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanOtherMessages Pass True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
func (void *VoidResponse) SetCanSendOtherMessage(can bool) *VoidResponse {
	body := JSON{
		"can_send_other_messages": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanAddWebPagePreview Pass True, if the user may add web page previews to their messages, implies can_send_media_messages
func (void *VoidResponse) SetCanAddWebPagePreview(can bool) *VoidResponse {
	body := JSON{
		"can_add_web_page_previews": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
PromoteChatMember Use this method to promote or demote a user in a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Pass False for all boolean parameters to demote a user. Returns True on success.
*/
func (client *Client) PromoteChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}

	url := client.baseURL + fmt.Sprintf(EndpointPromoteChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetCanChangeInfo Pass True, if the administrator can change chat title, photo and other settings
func (void *VoidResponse) SetCanChangeInfo(can bool) *VoidResponse {
	body := JSON{
		"can_change_info": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanPostMessage Pass True, if the administrator can create channel posts, channels only
func (void *VoidResponse) SetCanPostMessage(can bool) *VoidResponse {
	body := JSON{
		"can_post_messages": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanEditMessage Pass True, if the administrator can edit messages of other users and can pin messages, channels only
func (void *VoidResponse) SetCanEditMessage(can bool) *VoidResponse {
	body := JSON{
		"can_edit_messages": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanDeleteMessage Pass True, if the administrator can delete messages of other users
func (void *VoidResponse) SetCanDeleteMessage(can bool) *VoidResponse {
	body := JSON{
		"can_delete_messages": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanInviteUser Pass True, if the administrator can invite new users to the chat
func (void *VoidResponse) SetCanInviteUser(can bool) *VoidResponse {
	body := JSON{
		"can_invite_users": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanRestrictMember Pass True, if the administrator can restrict, ban or unban chat members
func (void *VoidResponse) SetCanRestrictMember(can bool) *VoidResponse {
	body := JSON{
		"can_restrict_members": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanPinMessage Pass True, if the administrator can pin messages, supergroups only
func (void *VoidResponse) SetCanPinMessage(can bool) *VoidResponse {
	body := JSON{
		"can_pin_messages": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetCanPromoteMember Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted,
// directly or indirectly (promoted by administrators that were appointed by him)
func (void *VoidResponse) SetCanPromoteMember(can bool) *VoidResponse {
	body := JSON{
		"can_promote_members": can,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
GetContent function for download file from telegram server, file path obtained from function GetFile()
Exp https://api.telegram.org/file/bot<token>/<file_path>
*/
func (client *Client) GetContent(path string) *VoidResponse {
	url := client.baseURL + fmt.Sprintf(EndpointGetContent, client.accessToken, path)
	request := gorequest.New().Get(url).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
UnbanChatMember Use this method to unban a previously kicked user in a supergroup or channel.
The user will not return to the group or channel automatically, but will be able to join via link, etc.
The bot must be an administrator for this to work. Returns True on success.
*/
func (client *Client) UnbanChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}

	url := client.baseURL + fmt.Sprintf(EndpointUnbanChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(url).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (void *VoidResponse) Commit() ([]byte, *http.Response, error) {
	var body []byte
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *bool `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, body, errs = void.Request.EndBytes()
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, void.Client.expBackOff); err != nil {
		return nil, MakeHTTPResponse(void.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return body, res, nil
}
