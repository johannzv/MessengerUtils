package models

// Message holds message
type Message struct {
	Mid         string      `json:"mid,omitempty"`
	Text        string      `json:"text"`
	QuickReplay QuickReplay `json:"quick_reply,omitempty"`
}

// Attachment struct holding attachments
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

// Payload may be multimedia or a location
type Payload struct {
	URL             string    `json:"url,omitempty"`
	Location        Location  `json:"location,omitempty"`
	TemplateType    string    `json:"template_type,omitempty"`
	Sharable        string    `json:"sharable,omitempty"`
	TopElementStyle string    `json:"top_element_style,omitempty"`
	Elements        []Element `json:"elements,omitempty"`
	Buttons         []Button  `json:"buttons,omitempty"`
}

//Button dfpsd
type Button struct {
	Title   string `json:"title,omitempty"`
	Type    string `json:"type,omitempty"`
	Payload string `json:"payload,omitempty"`
}

//Element of lsit
type Element struct {
	Title    string   `json:"title,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
	ImageURL string   `json:"image_url,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

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
