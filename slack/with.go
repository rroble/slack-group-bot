package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func hanleWith(c *gin.Context) {
	var req with

	if err := c.ShouldBindWith(&req, binding.FormPost); err != nil {
		fmt.Println("Bind error:", err)
		c.JSON(500, err)
		return
	}

	if !req.Allowed() {
		c.JSON(200, gin.H{
			"text": "Only allowed in private group",
		})
		return
	}

	fmt.Printf("%+v\n", req)

	users := getUsers(req.Text)
	if len(users) == 0 {
		c.Data(200, "text/plain", nil)
		return
	}

	// return asap
	c.JSON(200, gin.H{
		"text":          "",
		"response_type": "in_channel",
	})

	for _, user := range users {
		userChannel := gc.add(req.ChannelID, user)
		if userChannel != "" {
			postMessageToSlack(req.Text, userChannel, req.UserID, "added to group chat")
		}
	}
}

// ExtractUsers parse message and return found users' ID
func getUsers(message string) []string {
	start, end := false, false
	return strings.FieldsFunc(message, func(r rune) bool {
		switch r {
		case '@':
			start, end = true, false
			return true
		case '|':
			start, end = false, true
		}
		if !start || end {
			return true
		}
		return false
	})
}
