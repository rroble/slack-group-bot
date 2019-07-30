package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var gc *groupChat

func init() {
	gc = &groupChat{
		Users:  make(map[string]string, 0),
		Groups: make(map[string]users, 0),
	}
}

func hanleEvents(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Fprintf(c.Writer, "Bind error: %+v\n", err)
		c.JSON(500, err)
		return
	}

	fmt.Println(req)

	if req.IsVerification() {
		c.Data(200, "text/plain", []byte(req.Challenge))
		return
	}

	// ACK immediately
	c.Data(200, "text/plain", nil)

	// this is our bot's message, don't check
	if req.Event.ClientMsgID == "" {
		return
	}

	if req.Event.Type == "message" && req.Event.Text != "" {
		wrap := strings.Repeat("-", len(req.Event.Text)+2)
		fmt.Printf("%s\n|%s|\n%s\n", wrap, req.Event.Text, wrap)

		// 1. send to included users except the user himself
		// the current user is not possibly on the list
		if userChannels, ok := gc.Groups[req.Event.Channel]; ok {
			for uc, active := range userChannels {
				if active {
					postMessageToSlack(req.Event.Text, uc, req.Event.User, "group chat")
				}
			}
			return
		}

		// 2. or reply to group the user is included to
		// and only if sending from the user's channel
		// TODO: reply to specific group ie. @1 <msg>
		if userChannel, ok := gc.Users[req.Event.User]; ok && userChannel == req.Event.Channel {
			for groupChannel, userChannels := range gc.Groups {
				for uc, active := range userChannels {
					if uc == userChannel && active {
						postMessageToSlack(req.Event.Text, groupChannel, req.Event.User, "included in group chat")
						break // next group
					}
				}
			}
		}
	}

}
