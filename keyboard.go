package telegraph

type (
	InlineKeyboardButton struct {
		Text              string        `json:"text"`
		URL               string        `json:"url,omitempty"`
		CallbackData      string        `json:"callback_data,omitempty"`
		SwitchInlineQuery string        `json:"switch_inline_query,omitempty"`
		SwitchInlineChat  string        `json:"switch_inline_query_current_chat,omitempty"`
		Pay               bool          `json:"pay,omitempty"`
		CallbackGame      *CallbackGame `json:"callback_game,omitempty"`
	}

	ReplyKeyboardMarkup struct {
		Keyboard        [][]KeyboardButton `json:"keyboard"`
		ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
		OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
		Selective       bool               `json:"selective,omitempty"`
	}

	ReplyKeyboardRemove struct {
		RemoveKeyboard bool `json:"remove_keyboard"`
		Selective      bool `json:"selective,omitempty"`
	}

	KeyboardButton struct {
		Text            string `json:"text"`
		RequestContact  bool   `json:"request_contact,omitempty"`
		RequestLocation bool   `json:"request_location,omitempty"`
	}
)

/*
Requests clients to remove the custom keyboard (user will not be able to summon this keyboard;
if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)

Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1)
users that are @mentioned in the text of the Message object; 2)
if the bot's message is a reply (has reply_to_message_id), sender of the original message.

Example: A user votes in a poll,
bot returns confirmation message in reply to the vote and removes the keyboard for that user,
while still showing the keyboard with poll options to users who haven't voted yet.
*/
func NewReplyKeyboardRemove(selective bool) *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// NewReplyKeyboardMarkup Array of button rows, each represented by an Array of KeyboardButton objects
func NewReplyKeyboardMarkup(keyboards [][]KeyboardButton) *ReplyKeyboardMarkup {
	return &ReplyKeyboardMarkup{
		Keyboard: keyboards,
	}
}

// SetResizeKeyboard Optional. Requests clients to resize the keyboard vertically for optimal fit
// (e.g., make the keyboard smaller if there are just two rows of buttons).
// Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
func (reply *ReplyKeyboardMarkup) SetResizeKeyboard(resize bool) *ReplyKeyboardMarkup {
	reply.ResizeKeyboard = resize
	return reply
}

// SetOneTimeKeyboard Optional. Requests clients to hide the keyboard as soon as it's been used.
// The keyboard will still be available, but clients will automatically display the usual
// letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again.
// Defaults to false.
func (reply *ReplyKeyboardMarkup) SetOneTimeKeyboard(keyboard bool) *ReplyKeyboardMarkup {
	reply.OneTimeKeyboard = keyboard
	return reply
}

/*
SetSelective Optional. Use this parameter if you want to show the keyboard to specific users only.
Targets: 1) users that are @mentioned in the text of the Message object;
2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.

Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language.
Other users in the group don’t see the keyboard.
*/
func (reply *ReplyKeyboardMarkup) SetSelective(selective bool) *ReplyKeyboardMarkup {
	reply.Selective = selective
	return reply
}

// NewKeyboardButton This object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields are mutually exclusive.
func NewKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{
		Text: text,
	}
}

// SetRequestContact Optional. If True, the user's phone number will be sent as a contact when the button is pressed.
// Available in private chats only
func (keyboard *KeyboardButton) SetRequestContact(contact bool) *KeyboardButton {
	keyboard.RequestContact = contact
	return keyboard
}

// SetRequestLocation Optional. If True, the user's current location will be sent when the button is pressed.
// Available in private chats only
func (keyboard *KeyboardButton) SetRequestLocation(contact bool) *KeyboardButton {
	keyboard.RequestLocation = contact
	return keyboard
}

/*
NewInlineKeyboardButton This object represents one button of an inline keyboard.
You must use exactly one of the optional fields.

Label text on the button
*/
func NewInlineKeyboardButton(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{
		Text: text,
	}
}

// SetURL Optional. HTTP url to be opened when button is pressed
func (keyboard *InlineKeyboardButton) SetURL(url string) *InlineKeyboardButton {
	keyboard.URL = url
	return keyboard
}

// SetCallbackData Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
func (keyboard *InlineKeyboardButton) SetCallbackData(callback string) *InlineKeyboardButton {
	keyboard.CallbackData = callback
	return keyboard
}

/*
SetSwitchInlineQuery Optional. If set, pressing the button will prompt the user to select one of their chats,
open that chat and insert the bot‘s username and the specified inline query in the input field.
Can be empty, in which case just the bot’s username will be inserted.

Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it.
Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from,
skipping the chat selection screen.
*/
func (keyboard *InlineKeyboardButton) SetSwitchInlineQuery(query string) *InlineKeyboardButton {
	keyboard.SwitchInlineQuery = query
	return keyboard
}

/*
SetSwitchInlineChat Optional. If set,
pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field.
Can be empty, in which case only the bot’s username will be inserted.

This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
*/
func (keyboard *InlineKeyboardButton) SetSwitchInlineChat(chat string) *InlineKeyboardButton {
	keyboard.SwitchInlineChat = chat
	return keyboard
}

// SetPay Optional. Specify True, to send a Pay button.
// NOTE: This type of button must always be the first button in the first row.
func (keyboard *InlineKeyboardButton) SetPay(pay bool) *InlineKeyboardButton {
	keyboard.Pay = pay
	return keyboard
}

// SetCallbackGame Optional. Description of the game that will be launched when the user presses the button.
// NOTE: This type of button must always be the first button in the first row.
func (keyboard *InlineKeyboardButton) SetCallbackGame(callback *CallbackGame) *InlineKeyboardButton {
	keyboard.CallbackGame = callback
	return keyboard
}
