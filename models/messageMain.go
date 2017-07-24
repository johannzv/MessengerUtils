package models

//MessageObject  main holder of the event
type MessageObject struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

// Entry holds each entry
type Entry struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

// Messaging array of message objects
type Messaging struct {
	Sender    Agent    `json:"sender"`
	Recipient Agent    `json:"recipient"`
	Timestamp int64    `json:"timestamp"`
	Message   Message  `json:"message,omitempty"`
	Postback  Postback `json:"postback,omitempty"`
	Delivery  Delivery `json:"delivery,omitempty"`
	Read      Read     `json:"read,omitempty"`
}

// Agent sender or recieveier
type Agent struct {
	ID string `json:"id"`
}
