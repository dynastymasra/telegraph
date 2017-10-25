package telegraph

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup               = "group"
	ChatTypeSuperGroup          = "supergroup"
	ChatTypeChannel             = "channel"
)

type (
	Update struct {
		UpdateID          int64    `json:"update_id"`
		Message           *Message `json:"message,omitempty"`
		EditedMessage     *Message `json:"edited_message,omitempty"`
		ChannelPost       *Message `json:"channel_post,omitempty"`
		EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	}

	Message struct {
		MessageID int64  `json:"message_id"`
		Date      int64  `json:"date"`
		Chat      Chat   `json:"chat"`
		From      *From  `json:"from,omitempty"`
		Text      string `json:"text,omitempty"`
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
