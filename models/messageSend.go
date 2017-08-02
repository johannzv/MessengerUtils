package models

// MessageSend : the message that should be sent to a user
type MessageSend struct {
	Recipient    Agent              `json:"recipient"`
	Message      MessageSendingType `json:"message,omitempty"`
	SenderAction string             `json:"sender_action,omitempty"`
}

//MessageSendingType sendingMessage
type MessageSendingType struct {
	Text         string        `json:"text"`
	Attachment   Attachment    `json:"attachment,omitempty"`
	QuickReplies []QuickReplay `json:"quick_replies,omitempty"`
}
