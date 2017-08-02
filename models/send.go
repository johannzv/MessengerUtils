package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//MessageModel for sending a message
type MessageModel interface {
	Send(token string) (*http.Response, error)
}

// Send the message to a page given a valid token
func (msg MessageSend) Send(token string) (*http.Response, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err2 := http.Post("https://graph.facebook.com/v2.6/me/messages?access_token="+token, "application/json", body)
	fmt.Println(err2)
	return res, err2
}

func (msg MessageSendTopicMenu) Send(token string) (*http.Response, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err2 := http.Post("https://graph.facebook.com/v2.6/me/messages?access_token="+token, "application/json", body)
	fmt.Println(err2)

	return res, err2
}
