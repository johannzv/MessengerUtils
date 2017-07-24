package models

// Message holds message
type Message struct {
	Mid         string       `json:"mid"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments,omitempty"`
	QuickReplay QuickReplay  `json:"quick_reply, omitempty"`
}

// Attachment struct holding attachments
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

// Payload may be multimedia or a location
type Payload struct {
	URL      string   `json:"url,omitempty"`
	Location Location `json:"location,omitempty"`
}

// TODO test this

// Location cordinates
type Location struct {
	Lat  int64 `json:"coridantes.lat"`
	Long int64 `json:"coridantes.long"`
}

//QuickReplay :text provided via QuickReplay
type QuickReplay struct {
	Payload string `json:"payload"`
}
