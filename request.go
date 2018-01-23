package telegraph

import (
	"fmt"

	"net/http"
	"net/url"

	"github.com/cenkalti/backoff"
	"github.com/parnurzeal/gorequest"
)

type (
	// VoidResponse struct to handle request and response telegram api
	VoidResponse struct {
		Client  *Client
		Request *gorequest.SuperAgent
	}

	// StringResponse struct to handle request and response telegram api
	StringResponse struct {
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
SetWebHook Use this method to specify a url and receive incoming updates via an outgoing webhook.
Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
containing a JSON-serialized Update. In case of an unsuccessful request,
we will give up after a reasonable amount of attempts. Returns true.
+ webHook - HTTPS url to send updates to. Use an empty string to remove webhook integration

Available method can used with this method
+ SetCertificate()
+ SetMaxConnection()
+ SetAllowedUpdates()
*/
func (client *Client) SetWebHook(webHook string) *VoidResponse {
	body := JSON{
		"url": webHook,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSetWebHook, client.accessToken)
	request := gorequest.New().Post(endpoint).Type(gorequest.TypeJSON).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

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

/*
DeleteWebHook Use this method to remove webhook integration if you decide to switch back to getUpdates.
Returns True on success. Requires no parameters.
*/
func (client *Client) DeleteWebHook() *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointDeleteWebHook, client.accessToken)
	request := gorequest.New().Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
EditMessageLiveLocation Use this method to edit live location messages sent by the bot or via the bot (for inline bots).
A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
+ latitude - Latitude of new location
+ longitude - Longitude of new location

Available method can used with this method
+ SetChatID()
+ SetMessageID()
+ SetInlineMessageID()
+ SetReplyMarkup()
*/
func (client *Client) EditMessageLiveLocation(latitude, longitude float64) *VoidResponse {
	body := JSON{
		"latitude":  latitude,
		"longitude": longitude,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointEditMessageLiveLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
StopMessageLiveLocation Use this method to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires.
On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.

Available method can used with this method
+ SetChatID()
+ SetMessageID()
+ SetInlineMessageID()
+ SetReplyMarkup()
*/
func (client *Client) StopMessageLiveLocation() *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointStopMessageLiveLocation, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
EditMessageText Use this method to edit text and game messages sent by the bot or via the bot (for inline bots).
On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
+ text - New text of the message

Available method can used with this method
+ SetChatID()
+ SetMessageID()
+ SetInlineMessageID()
+ SetParseMode()
+ SetDisableWebPagePreview()
+ SetReplyMarkup()
*/
func (client *Client) EditMessageText(text string) *VoidResponse {
	body := JSON{
		"text": text,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointEditMessageText, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetParseMode Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
func (void *VoidResponse) SetParseMode(mode string) *VoidResponse {
	body := JSON{
		"parse_mode": mode,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetDisableWebPagePreview Disables link previews for links in this message
func (void *VoidResponse) SetDisableWebPagePreview(disable bool) *VoidResponse {
	body := JSON{
		"disable_web_page_preview": disable,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
EditMessageCaption Use this method to edit captions of messages sent by the bot or via the bot (for inline bots).
On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
+ caption - New caption of the message

Available method can used with this method
+ SetChatID()
+ SetMessageID()
+ SetInlineMessageID()
+ SetReplyMarkup()
*/
func (client *Client) EditMessageCaption(caption string) *VoidResponse {
	body := JSON{
		"caption": caption,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointEditMessageCaption, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
EditMessageReplyMarkup Use this method to edit only the reply markup of messages sent by the bot or via the bot (for inline bots).
On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.

Available method can used with this method
+ SetChatID()
+ SetMessageID()
+ SetInlineMessageID()
+ SetReplyMarkup()
*/
func (client *Client) EditMessageReplyMarkup() *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointEditMessageReplyMarkup, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version)

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

// SetReplyMarkup A JSON-serialized object for a new inline keyboard.
func (void *VoidResponse) SetReplyMarkup(inline [][]InlineKeyboardButton) *VoidResponse {
	body := JSON{
		"reply_markup": JSON{
			"inline_keyboard": inline,
		},
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
DeleteMessage Use this method to delete a message, including service messages, with the following limitations:
- A message can only be deleted if it was sent less than 48 hours ago.
- Bots can delete outgoing messages in groups and supergroups.
- Bots granted can_post_messages permissions can delete outgoing messages in channels.
- If the bot is an administrator of a group, it can delete any message there.
- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ messageId - Identifier of the message to delete
*/
func (client *Client) DeleteMessage(chatId interface{}, messageId int64) *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointDeleteMessage, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v&message_id=%v", chatId, messageId))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SendChatAction Use this method when you need to tell the user that something is happening on the bot's side.
The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
Returns True on success.

Example: The ImageBot needs some time to process a request and upload the image.
Instead of sending a text message along the lines of “Retrieving image, please wait…”,
the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ action - Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages,
  upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files,
  upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
*/
func (client *Client) SendChatAction(chatId interface{}, action string) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"action":  action,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSendChatAction, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

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
+ chatId - Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
+ userId - Unique identifier of the target user

Available method can used with this method
+ SetUntilDate()
*/
func (client *Client) KickChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointKickChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
RestrictChatMember Use this method to restrict a user in a supergroup.
The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
+ userId - Unique identifier of the target user

Available method can used with this method
+ SetUntilDate()
+ SetCanSendMessage()
+ SetCanSendMediaMessage()
+ SetCanSendOtherMessage()
+ SetCanAddWebPagePreview()
*/
func (client *Client) RestrictChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointRestrictChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

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
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ userId - Unique identifier of the target user

Available method can used with this method
+ SetCanChangeInfo()
+ SetCanPostMessage()
+ SetCanEditMessage()
+ SetCanDeleteMessage()
+ SetCanInviteUser()
+ SetCanRestrictMember()
+ SetCanPinMessage()
+ SetCanPromoteMember()
*/
func (client *Client) PromoteChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointPromoteChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

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
	endpoint := client.baseURL + fmt.Sprintf(EndpointGetContent, client.accessToken, path)
	request := gorequest.New().Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
UnbanChatMember Use this method to unban a previously kicked user in a supergroup or channel.
The user will not return to the group or channel automatically, but will be able to join via link, etc.
The bot must be an administrator for this to work. Returns True on success.
+ chatId - Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
+ userId - Unique identifier of the target user
*/
func (client *Client) UnbanChatMember(chatId interface{}, userId int64) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"user_id": userId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointUnbanChatMember, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatPhoto Use this method to set a new profile photo for the chat. Photos can't be changed for private chats.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
Note: In regular groups (non-supergroups), this method will only work if the ‘All Members Are Admins’ setting is off in the target group.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ photo - New chat photo, uploaded using multipart/form-data
*/
func (client *Client) SetChatPhoto(chatId interface{}, photo string) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSetChatPhoto, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeMultipart).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Send(body).SendFile(photo, "", "photo")

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
DeleteChatPhoto Use this method to delete a chat photo. Photos can't be changed for private chats.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
Note: In regular groups (non-supergroups), this method will only work if the ‘All Members Are Admins’ setting is off in the target group.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
*/
func (client *Client) DeleteChatPhoto(chatId interface{}) *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointDeleteChatPhoto, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatId))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatTitle Use this method to change the title of a chat. Titles can't be changed for private chats.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
Note: In regular groups (non-supergroups), this method will only work if the ‘All Members Are Admins’ setting is off in the target group.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ title - New chat title, 1-255 characters
*/
func (client *Client) SetChatTitle(chatId interface{}, title string) *VoidResponse {
	body := JSON{
		"chat_id": chatId,
		"title":   title,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSetChatTitle, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatDescription Use this method to change the description of a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ description - New chat description, 0-255 characters
*/
func (client *Client) SetChatDescription(chatId interface{}, description string) *VoidResponse {
	body := JSON{
		"chat_id":     chatId,
		"description": description,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSetChatDescription, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
PinChatMessage Use this method to pin a message in a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel.
Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
+ messageId - Identifier of a message to pin

Available method can used with this method
+ SetDisableNotification()
*/
func (client *Client) PinChatMessage(chatId interface{}, messageId int64) *VoidResponse {
	body := JSON{
		"chat_id":    chatId,
		"message_id": messageId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointPinChatMessage, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetDisableNotification Pass True, if it is not necessary to send a notification to all chat members about the new pinned message.
// Notifications are always disabled in channels.
func (void *VoidResponse) SetDisableNotification(disable bool) *VoidResponse {
	body := JSON{
		"disable_notification": disable,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
UnpinChatMessage Use this method to unpin a message in a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel.
Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
*/
func (client *Client) UnpinChatMessage(chatId interface{}) *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointUnpinChatMessage, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatId))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
LeaveChat Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
*/
func (client *Client) LeaveChat(chatId interface{}) *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointLeaveChat, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatId))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
SetChatStickerSet Use this method to set a new group sticker set for a supergroup.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
+ name - Name of the sticker set to be set as the group sticker set
*/
func (client *Client) SetChatStickerSet(chatId interface{}, name string) *VoidResponse {
	body := JSON{
		"chat_id":          chatId,
		"sticker_set_name": name,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSetChatStickerSet, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
Returns True on success.
+ chatId - Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
*/
func (client *Client) DeleteChatStickerSet(chatId interface{}) *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointDeleteChatStickerSet, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatId))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
AnswerCallbackQuery Use this method to send answers to callback queries sent from inline keyboards.
The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.

Alternatively, the user can be redirected to the specified Game URL.
For this option to work, you must first create a game for your bot via @Botfather and accept the terms.
Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
+ queryId - Unique identifier for the query to be answered

Available method can used with this method
+ SetText()
+ SetShowAlert()
+ SetURL()
+ SetCacheTime()
*/
func (client *Client) AnswerCallbackQuery(queryId string) *VoidResponse {
	body := JSON{
		"callback_query_id": queryId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointAnswerCallbackQuery, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetText Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
func (void *VoidResponse) SetText(text string) *VoidResponse {
	body := JSON{
		"text": text,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetShowAlert If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
func (void *VoidResponse) SetShowAlert(show bool) *VoidResponse {
	body := JSON{
		"show_alert": show,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetURL URL that will be opened by the user's client.
// If you have created a Game and accepted the conditions via @Botfather,
// specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.
func (void *VoidResponse) SetURL(url string) *VoidResponse {
	body := JSON{
		"url": url,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
CreateNewStickerSet Use this method to create new sticker set owned by a user.
The bot will be able to edit the created sticker set. Returns True on success.
- userId User identifier of created sticker set owner
- name Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals).
  Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”.
  <bot_username> is case insensitive. 1-64 characters
- title Sticker set title, 1-64 characters
- pngSticker Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px.
  Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data
- emojis One or more emoji corresponding to the sticker
- upload set true if upload file to telegram from local

+ userId - User identifier of created sticker set owner
+ name - Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores.
  Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”.
  <bot_username> is case insensitive. 1-64 characters.
+ title - Sticker set title, 1-64 characters
+ pngSticker - Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px.
  Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
+ emojis - One or more emoji corresponding to the sticker

Available method can used with this method
+ SetContainsMask()
+ SetMaskPosition()
*/
func (client *Client) CreateNewStickerSet(userId int64, name, title, pngSticker, emojis string) *VoidResponse {
	body := JSON{
		"user_id":     userId,
		"name":        name,
		"title":       title,
		"png_sticker": pngSticker,
		"emojis":      emojis,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointCreateNewStickerSet, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(pngSticker); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(pngSticker, "", "png_sticker")
	}

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetContainsMask Pass True, if a set of mask stickers should be created
func (void *VoidResponse) SetContainsMask(mask bool) *VoidResponse {
	body := JSON{
		"contains_masks": mask,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
AddStickerToSet Use this method to add a new sticker to a set created by the bot. Returns True on success.
- userId User identifier of sticker set owner
- name Sticker set name
- pngSticker Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px.
  Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
- emojis One or more emoji corresponding to the sticker
- upload set true if upload file to telegram from local

+ userId - User identifier of sticker set owner
+ name - Sticker set name
+ pngSticker - Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px.
  Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
+ emojis - One or more emoji corresponding to the sticker

Available method can used with this method
+ SetMaskPosition()
*/
func (client *Client) AddStickerToSet(userId int64, name, pngSticker, emojis string) *VoidResponse {
	body := JSON{
		"user_id":     userId,
		"name":        name,
		"png_sticker": pngSticker,
		"emojis":      emojis,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointAddStickerToSet, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	if _, err := url.ParseRequestURI(pngSticker); err != nil {
		request.Type(gorequest.TypeMultipart).SendFile(pngSticker, "", "png_sticker")
	}

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetMaskPosition A JSON-serialized object for position where the mask should be placed on faces
func (void *VoidResponse) SetMaskPosition(mask MaskPosition) *VoidResponse {
	body := JSON{
		"mask_position": mask,
	}
	void.Request = void.Request.Send(body)

	return void
}

/*
SetStickerPositionInSet Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.
- sticker File identifier of the sticker
- position New sticker position in the set, zero-based

+ sticker - File identifier of the sticker
+ position - New sticker position in the set, zero-based
*/
func (client *Client) SetStickerPositionInSet(sticker string, position int) *VoidResponse {
	body := JSON{
		"sticker":  sticker,
		"position": position,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointSetStickerPositionInSet, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Type(gorequest.TypeJSON).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
DeleteStickerFromSet Use this method to delete a sticker from a set created by the bot. Returns True on success.
+ sticker - File identifier of the sticker
*/
func (client *Client) DeleteStickerFromSet(sticker string) *VoidResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointDeleteStickerFromSet, client.accessToken)
	request := gorequest.New().Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Type(gorequest.TypeJSON).
		Query(fmt.Sprintf("sticker=%v", sticker))

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

/*
AnswerInlineQuery Use this method to send answers to an inline query. On success, True is returned.
No more than 50 results per query are allowed.
+ inlineQueryId - Unique identifier for the answered query
+ result - A JSON-serialized array of results for the inline query

Available method can used with this method
+ SetCacheTime()
+ SetIsPersonal()
+ SetNextOffset()
+ SetSwitchPMText()
+ SetSwitchPMParameter()
*/
func (client *Client) AnswerInlineQuery(inlineQueryId string, result ...JSON) *VoidResponse {
	body := JSON{
		"inline_query_id": inlineQueryId,
		"results":         result,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointAnswerInlineQuery, client.accessToken)
	request := gorequest.New().Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Type(gorequest.TypeJSON).Send(body)

	return &VoidResponse{
		Client:  client,
		Request: request,
	}
}

// SetCacheTime The maximum amount of time in seconds that the result of the callback query may be cached client-side.
// Telegram apps will support caching starting in version 3.14. Defaults to 0.
//
// The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
func (void *VoidResponse) SetCacheTime(time int64) *VoidResponse {
	body := JSON{
		"cache_time": time,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetIsPersonal Pass True, if results may be cached on the server side only for the user that sent the query.
// By default, results may be returned to any user who sends the same query
func (void *VoidResponse) SetIsPersonal(personal bool) *VoidResponse {
	body := JSON{
		"is_personal": personal,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetNextOffset Pass the offset that a client should send in the next query with the same text to receive more results.
// Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
func (void *VoidResponse) SetNextOffset(offset string) *VoidResponse {
	body := JSON{
		"next_offset": offset,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetSwitchPMText If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
func (void *VoidResponse) SetSwitchPMText(text string) *VoidResponse {
	body := JSON{
		"switch_pm_text": text,
	}
	void.Request = void.Request.Send(body)

	return void
}

// SetSwitchPMParameter Deep-linking parameter for the /start message sent to the bot when user presses the switch button.
// 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
// Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly.
// To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any.
// The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link.
// Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
func (void *VoidResponse) SetSwitchPMParameter(param string) *VoidResponse {
	body := JSON{
		"switch_pm_parameter": param,
	}
	void.Request = void.Request.Send(body)

	return void
}

// Commit execute request to telegram
func (void *VoidResponse) Commit() ([]byte, *http.Response, error) {
	var body []byte
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
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

/*
ExportChatInviteLink Use this method to export an invite link to a supergroup or a channel.
The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
Returns exported invite link as String on success.
+ chatId - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
*/
func (client *Client) ExportChatInviteLink(chatId interface{}) *StringResponse {
	body := JSON{
		"chat_id": chatId,
	}
	endpoint := client.baseURL + fmt.Sprintf(EndpointExportChatInviteLink, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Post(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).Send(body)

	return &StringResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (void *StringResponse) Commit() (string, *http.Response, error) {
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result string `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, _, errs = void.Request.EndStruct(&model)
		if len(errs) > 0 {
			return errs[0]
		}
		return nil
	}

	if err := backoff.Retry(operation, void.Client.expBackOff); err != nil {
		return "", MakeHTTPResponse(void.Request), err
	}
	if res.StatusCode != http.StatusOK {
		return "", res, fmt.Errorf("%v %v", model.ErrorCode, model.Description)
	}

	return model.Result, res, nil
}

/*
GetChatMembersCount Use this method to get the number of members in a chat. Returns Int on success.
+ chatId - Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
*/
func (client *Client) GetChatMembersCount(chatId interface{}) *IntegerResponse {
	endpoint := client.baseURL + fmt.Sprintf(EndpointGetChatMembersCount, client.accessToken)
	request := gorequest.New().Type(gorequest.TypeJSON).Get(endpoint).Set(UserAgentHeader, UserAgent+"/"+Version).
		Query(fmt.Sprintf("chat_id=%v", chatId))

	return &IntegerResponse{
		Client:  client,
		Request: request,
	}
}

// Commit execute request to telegram
func (void *IntegerResponse) Commit() (*int64, *http.Response, error) {
	var errs []error

	res := &http.Response{}
	model := struct {
		ErrorResponse
		Result *int64 `json:"result,omitempty"`
	}{}

	operation := func() error {
		res, _, errs = void.Request.EndStruct(&model)
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

	return model.Result, res, nil
}
