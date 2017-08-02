package models

// MessageSend : the message that should be sent to a user
type MessageSend struct {
	Recipient    Agent              `json:"recipient"`
	Message      MessageSendingType `json:"message,omitempty"`
	SenderAction string             `json:"sender_action,omitempty"`
}

//TODO
type MessageSendTopicMenu struct {
	Recipient Agent           `json:"recipient"`
	Message   MessageTemplate `json:"message,omitempty"`
}

//MessageSendingType sendingMessage
type MessageSendingType struct {
	Text         string        `json:"text,omitempty"`
	QuickReplies []QuickReplay `json:"quick_replies,omitempty"`
}

type MessageTemplate struct {
	Attachment Attachment `json:"attachment,omitempty"`
}
