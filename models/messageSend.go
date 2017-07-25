package models

// MessageSend : the message that should be sent to a user
type MessageSend struct {
	Recipient    Agent   `json:"recipient"`
	Message      Message `json:"message,omitempty"`
	SenderAction string  `json:"sender_action,omitempty"`
}
