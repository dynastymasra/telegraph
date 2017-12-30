package telegraph

import (
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type (
	ChatType   string
	StatusType string
	MediaType  string
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

	MediaTypePhoto MediaType = "photo"
	MediaTypeVideo           = "video"
)

type (
	// JSON struct json type
	JSON map[string]interface{}

	// ErrorResponse struct parse error response from telegram
	ErrorResponse struct {
		OK          bool   `json:"ok"`
		ErrorCode   int    `json:"error_code,omitempty"`
		Description string `json:"description,omitempty"`
	}

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

	// InlineQuery This object represents an incoming inline query. When the user sends an empty query,
	// your bot could return some default or trending results.
	InlineQuery struct {
		ID       string    `json:"id"`
		From     User      `json:"from"`
		Location *Location `json:"location,omitempty"`
		Query    string    `json:"query"`
		Offset   string    `json:"offset"`
	}

	// Location This object represents a point on the map.
	Location struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}

	// ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
	ChosenInlineResult struct {
		ResultID        string    `json:"result_id"`
		From            User      `json:"from"`
		Location        *Location `json:"location,omitempty"`
		InlineMessageID string    `json:"inline_message_id,omitempty"`
		Query           string    `json:"query"`
	}

	// CallbackQuery This object represents an incoming callback query from a callback button in an inline keyboard.
	// If the button that originated the query was attached to a message sent by the bot,
	// the field message will be present. If the button was attached to a message sent via the bot (in inline mode),
	// the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
	CallbackQuery struct {
		ID              string   `json:"id"`
		From            User     `json:"from"`
		Message         *Message `json:"message,omitempty"`
		InlineMessageID string   `json:"inline_message_id,omitempty"`
		ChatInstance    string   `json:"chat_instance"`
		Data            string   `json:"data,omitempty"`
		GameShortName   string   `json:"game_short_name,omitempty"`
	}

	// ShippingQuery This object contains information about an incoming shipping query.
	ShippingQuery struct {
		ID              string          `json:"id"`
		From            User            `json:"from"`
		InvoicePayload  string          `json:"invoice_payload"`
		ShippingAddress ShippingAddress `json:"shipping_address"`
	}

	// ShippingAddress This object represents a shipping address.
	ShippingAddress struct {
		CountryCode string `json:"country_code"`
		State       string `json:"state"`
		City        string `json:"city"`
		StreetLine1 string `json:"street_line1"`
		StreetLine2 string `json:"street_line2"`
		PostCode    string `json:"post_code"`
	}

	// PreCheckoutQuery This object contains information about an incoming pre-checkout query.
	PreCheckoutQuery struct {
		ID               string `json:"id"`
		From             User   `json:"from"`
		Currency         string `json:"currency"`
		TotalAmount      int64  `json:"total_amount"`
		InvoicePayload   string `json:"invoice_payload"`
		ShippingOptionID string `json:"shipping_option_id,omitempty"`
		OrderInfo        string `json:"order_info,omitempty"`
	}

	// Message This object represents a message.
	Message struct {
		MessageID             int64              `json:"message_id"`
		From                  *User              `json:"from,omitempty"`
		Date                  int64              `json:"date"`
		Chat                  Chat               `json:"chat"`
		ForwardFrom           *User              `json:"forward_from,omitempty"`
		ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`
		ForwardFromMessageID  int64              `json:"forward_from_message_id,omitempty"`
		ForwardSignature      string             `json:"forward_signature,omitempty"`
		ForwardDate           int64              `json:"forward_date,omitempty"`
		ReplyToMessage        string             `json:"reply_to_message,omitempty"`
		EditDate              int64              `json:"edit_date,omitempty"`
		AuthorSignature       string             `json:"author_signature,omitempty"`
		Text                  string             `json:"text,omitempty"`
		Entities              []MessageEntity    `json:"entities,omitempty"`
		CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
		Audio                 *Audio             `json:"audio,omitempty"`
		Document              *Document          `json:"document,omitempty"`
		Game                  *Game              `json:"game,omitempty"`
		Photos                []PhotoSize        `json:"photo,omitempty"`
		Sticker               *Sticker           `json:"sticker,omitempty"`
		Video                 *Video             `json:"video,omitempty"`
		Voice                 *Voice             `json:"voice,omitempty"`
		VideoNote             *VideoNote         `json:"video_note,omitempty"`
		Caption               string             `json:"caption,omitempty"`
		Contact               *Contact           `json:"contact,omitempty"`
		Location              *Location          `json:"location,omitempty"`
		Venue                 *Venue             `json:"venue,omitempty"`
		NewChatMembers        []User             `json:"new_chat_members,omitempty"`
		LeftChatMember        *User              `json:"left_chat_member,omitempty"`
		NewChatTitle          string             `json:"new_chat_title,omitempty"`
		NewChatPhoto          []PhotoSize        `json:"new_chat_photo,omitempty"`
		DeleteChatPhoto       bool               `json:"delete_chat_photo,omitempty"`
		GroupChatCreated      bool               `json:"group_chat_created,omitempty"`
		SuperGroupChatCreated bool               `json:"supergroup_chat_created,omitempty"`
		ChannelChatCreated    bool               `json:"channel_chat_created,omitempty"`
		MigrateToChatID       int64              `json:"migrate_to_chat_id,omitempty"`
		MigrateFromChatID     int64              `json:"migrate_from_chat_id,omitempty"`
		PinnedMessage         *Message           `json:"pinned_message,omitempty"`
		Invoice               *Invoice           `json:"invoice,omitempty"`
		SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`
	}

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

	// User This object represents a Telegram user or bot.
	User struct {
		ID           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name,omitempty"`
		Username     string `json:"username,omitempty"`
		LanguageCode string `json:"language_code,omitempty"`
	}

	// MessageEntity This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
	MessageEntity struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		URL    string `json:"url,omitempty"`
		User   *User  `json:"user,omitempty"`
	}

	// Audio This object represents an audio file to be treated as music by the Telegram clients
	Audio struct {
		FileID    string `json:"file_id"`
		Duration  int    `json:"duration"`
		Performer string `json:"performer,omitempty"`
		Title     string `json:"title,omitempty"`
		MimeType  string `json:"mime_type,omitempty"`
		FileSize  int    `json:"file_size,omitempty"`
	}

	// Document This object represents a general file (as opposed to photos, voice messages and audio files).
	Document struct {
		FileID   string     `json:"file_id"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		FileName string     `json:"file_name,omitempty"`
		MimeType string     `json:"mime_type,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
	}

	// PhotoSize This object represents one size of a photo or a file / sticker thumbnail.
	PhotoSize struct {
		FileID   string `json:"file_id"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		FileSize int    `json:"file_size"`
	}

	// Game This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
	Game struct {
		Title        string          `json:"title"`
		Description  string          `json:"description"`
		Photo        []PhotoSize     `json:"photo"`
		Text         string          `json:"text,omitempty"`
		TextEntities []MessageEntity `json:"text_entities,animation"`
		Animation    *Animation      `json:"animation,animation"`
	}

	// Animation You can provide an animation for your game so that it looks stylish in chats (check out Lumberjack for an example).
	// This object represents an animation file to be displayed in the message containing a game.
	Animation struct {
		FileID   string     `json:"file_id"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		FileName string     `json:"file_name,omitempty"`
		MimeType string     `json:"mime_type,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
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

	// MaskPosition This object describes the position on faces where a mask should be placed by default.
	MaskPosition struct {
		Point  string  `json:"point"`
		XShift float64 `json:"x_shift"`
		YShift float64 `json:"y_shift"`
		Scale  float64 `json:"scale"`
	}

	// Video This object represents a sticker.
	Video struct {
		FileID   string     `json:"file_id"`
		Width    int64      `json:"width"`
		Height   int64      `json:"height"`
		Duration int        `json:"duration"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		MimeType string     `json:"mime_type,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
	}

	// Voice This object represents a voice note.
	Voice struct {
		FileID   string `json:"file_id"`
		Duration int    `json:"duration"`
		MimeType string `json:"mime_type,omitempty"`
		FileSize int    `json:"file_size,omitempty"`
	}

	// VideoNote This object represents a video message (available in Telegram apps as of v.4.0).
	VideoNote struct {
		FileID   string     `json:"file_id"`
		Length   int        `json:"length"`
		Duration int        `json:"duration"`
		Thumb    *PhotoSize `json:"thumb,omitempty"`
		FileSize int        `json:"file_size,omitempty"`
	}

	// Contact This object represents a phone contact.
	Contact struct {
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name,omitempty"`
		UserID      int64  `json:"user_id,omitempty"`
	}

	// Venue This object represents a venue.
	Venue struct {
		Location     Location `json:"location"`
		Title        string   `json:"title"`
		Address      string   `json:"address"`
		FoursquareID string   `json:"foursquare_id,omitempty"`
	}

	// Invoice This object contains basic information about an invoice.
	Invoice struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		StartParameter string `json:"start_parameter"`
		Currency       string `json:"currency"`
		TotalAmount    int64  `json:"total_amount"`
	}

	// SuccessfulPayment This object contains basic information about a successful payment.
	SuccessfulPayment struct {
		Currency                string     `json:"currency"`
		TotalAmount             int64      `json:"total_amount"`
		InvoicePayload          string     `json:"invoice_payload"`
		ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
		OrderInfo               *OrderInfo `json:"order_info,omitempty"`
		TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
		ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
	}

	// OrderInfo This object represents information about an order.
	OrderInfo struct {
		Name            string           `json:"name,omitempty"`
		PhoneNumber     string           `json:"phone_number,omitempty"`
		Email           string           `json:"email,omitempty"`
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	}

	// WebhookInfo Contains information about the current status of a webhook.
	WebhookInfo struct {
		URL                  string   `json:"url"`
		HasCustomCertificate bool     `json:"has_custom_certificate"`
		PendingUpdateCount   int      `json:"pending_update_count"`
		LastErrorDate        int64    `json:"last_error_date,omitempty"`
		LastErrorMessage     string   `json:"last_error_message,omitempty"`
		MaxConnections       int      `json:"max_connections,omitempty"`
		AllowedUpdates       []string `json:"allowed_updates,omitempty"`
	}

	// Telegram clients will display a reply interface to the user
	// (act as if the user has selected the bot‘s message and tapped ’Reply').
	// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
	ForceReply struct {
		ForceReply bool `json:"force_reply"`
		Selective  bool `json:"selective,omitempty"`
	}

	// InlineKeyboardButton This object represents one button of an inline keyboard.
	// You must use exactly one of the optional fields.
	InlineKeyboardButton struct {
		Text                         string        `json:"text"`
		URL                          string        `json:"url,omitempty"`
		CallbackData                 string        `json:"callback_data,omitempty"`
		SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
		SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
		CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
		Pay                          bool          `json:"pay,omitempty"`
	}

	// CallbackGame A placeholder, currently holds no information. Use BotFather to set up your game.
	CallbackGame struct{}

	// ReplyKeyboardMarkup This object represents a custom keyboard with reply options
	// (see Introduction to bots for details and examples).
	ReplyKeyboardMarkup struct {
		Keyboard        [][]KeyboardButton `json:"keyboard"`
		ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
		OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
		Selective       bool               `json:"selective,omitempty"`
	}

	// KeyboardButton This object represents one button of the reply keyboard.
	// For simple text buttons String can be used instead of this object to specify text of the button.
	// Optional fields are mutually exclusive.
	KeyboardButton struct {
		Text            string `json:"text"`
		RequestContact  bool   `json:"request_contact,omitempty"`
		RequestLocation bool   `json:"request_location,omitempty"`
	}

	// ReplyKeyboardRemove Upon receiving a message with this object,
	// Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
	// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
	// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
	ReplyKeyboardRemove struct {
		RemoveKeyboard bool `json:"remove_keyboard"`
		Selective      bool `json:"selective,omitempty"`
	}

	// InputMedia This object represents the content of a media message to be sent. It should be one of
	// InputMediaPhoto
	// InputMediaVideo
	// See documentation for details https://core.telegram.org/bots/api#inputmedia
	InputMedia struct {
		Type     MediaType `json:"type"`
		Media    string    `json:"media"`
		Caption  string    `json:"caption,omitempty"`
		Width    int       `json:"width,omitempty"`
		Height   int       `json:"height,omitempty"`
		Duration int       `json:"duration,omitempty"`
	}
)

// MakeHTTPResponse create mock http response if request to API is error internal
func MakeHTTPResponse(agent *gorequest.SuperAgent) *http.Response {
	request, err := agent.MakeRequest()
	if err != nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}
	}

	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     request.Header,
		Request:    request,
	}
}
