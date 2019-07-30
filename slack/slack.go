package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var slackToken string

func init() {
	slackToken = os.Getenv("SLACK_API_TOKEN")
}

func postMessageToSlack(message, channel, sender string, attachments ...string) error {
	msg := gin.H{
		"text":     message,
		"channel":  channel,
		"as_user":  true,
		"username": sender,
	}
	if len(attachments) > 0 {
		atts := make([]gin.H, len(attachments))
		for i, att := range attachments {
			atts[i] = gin.H{
				"text": att,
			}
		}
		msg["attachments"] = atts
	}
	r, _ := json.Marshal(msg)
	url := "https://slack.com/api/chat.postMessage"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(r))
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+slackToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Error response:", body)
	}

	return err
}

func getUserChannel(user string) (string, error) {
	m := gin.H{
		"token": slackToken,
		"user":  user,
	}
	r, _ := json.Marshal(m)
	url := "https://slack.com/api/im.open"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(r))
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+slackToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	var chat im
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &chat); err != nil {
		return "", err
	}

	return chat.Channel.ID, nil
}
