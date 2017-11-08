package telegraph

type (
	ForwardMessage struct {
		ChatID              string `json:"chat_id"`
		FromChatID          string `json:"from_chat_id"`
		DisableNotification bool   `json:"disable_notification,omitempty"`
		MessageID           int64  `json:"message_id"`
		endpoint            string `json:"-"`
	}
)

// NewForwardMessage Use this method to forward messages of any kind. On success, the sent Message is returned.
func NewForwardMessage(chatID, fromChatID string, messageID int64) *ForwardMessage {
	return &ForwardMessage{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
		endpoint:   EndpointForwardMessage,
	}
}

// SetDisableNotification Sends the message silently. Users will receive a notification with no sound.
func (forward *ForwardMessage) SetDisableNotification(disable bool) *ForwardMessage {
	forward.DisableNotification = disable
	return forward
}
