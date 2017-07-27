package models

// Message holds message
type Message struct {
	Mid         string        `json:"mid,omitempty"`
	Text        string        `json:"text"`
	Attachments []Attachment  `json:"attachments,omitempty"`
	QuickReplay []QuickReplay `json:"quick_reply,omitempty"`
}

// Attachment struct holding attachments
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

// Payload may be multimedia or a location
type Payload struct {
	URL          string   `json:"url,omitempty"`
	Location     Location `json:"location,omitempty"`
	TemplateType string   `json:"template_type,omitempty"`
}

// TODO test this

// Location cordinates
type Location struct {
	Lat  int64 `json:"coridantes.lat"`
	Long int64 `json:"coridantes.long"`
}

//QuickReplay :text provided via QuickReplay
type QuickReplay struct {
	Payload     string `json:"payload,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	Title       string `json:"title,omitempty"`
}
