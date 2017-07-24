package models

//Postback struct
type Postback struct {
	Payload  string   `json:"payload"`
	Referral Referral `json:"referral,omitempty"`
}

//Delivery callback when a message is delivered
type Delivery struct {
	Mids      []string `json:"mids"`
	Watermark int64    `json:"watermark"`
	Seq       int32    `json:"seq"`
}

// Read callback when a message is read
type Read struct {
	Watermark int64 `json:"watermark"`
	Seq       int32 `json:"seq"`
}

// Referral :This callback will occur when the user already has a thread with the bot and user comes to the thread from: ad, link, scanning
type Referral struct {
	Ref    string `json:"ref,omitempty"`
	Source string `json:"source,omitempty"`
	Type   string `json:"type"`
	ADID   string `json:"ad_id,omitempty"`
}
