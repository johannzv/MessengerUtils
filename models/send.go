package models

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//MessageModel for sending a message
type MessageModel interface {
	send(token string) (*http.Response, error)
}

// sends the message to a page given a valid token
func (msg MessageSend) send(token string) (*http.Response, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return nil, err
	}
	res, err2 := http.Post("https://graph.facebook.com/v2.6/me/messages?access_token="+token, "application/json", body)
	return res, err2
}
