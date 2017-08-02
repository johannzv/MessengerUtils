package models

import (
	"encoding/json"
	"net/http"
	"time"
)

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

//UserMinimal : minimal amount of userdate to be stored on a user
type UserMinimal struct {
	PSID      string    `json:"PSID"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Locale    string    `json:"locale,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}

type ExtraInfo struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Locale    string `json:"locale,omitempty"`
}

//UserInfo interface for fetching parts
type UserInfo interface {
	Find(token string) error
}

//Find sumplements the extra fileds needed for minimal user
func (user *UserMinimal) Find(token string) error {
	var extrainfo ExtraInfo
	res, err := http.Get("https://graph.facebook.com/v2.6/" + user.PSID + "?fields=first_name,last_name,locale&access_token=" + token)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&extrainfo)
	if err != nil {
		return err
	}
	user.FirstName = extrainfo.FirstName
	user.LastName = extrainfo.LastName
	user.Locale = extrainfo.Locale
	return err
}
