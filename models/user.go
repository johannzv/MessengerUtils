package models

import "time"

// User holds the user object
type User struct {
	PSID       string    `json:"PSID"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	ProfilePic string    `json:"profilePic"`
	TimeZone   int       `json:"timeZone"`
	Gender     string    `json:"gender"`
	Created    time.Time `json:"created,omitempty"`
}
