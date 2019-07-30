package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func hanleWithout(c *gin.Context) {
	var req with

	if err := c.ShouldBindWith(&req, binding.FormPost); err != nil {
		fmt.Println("Bind error:", err)
		c.JSON(500, err)
		return
	}

	fmt.Printf("req: %+v\n", req)

	users := getUsers(req.Text)
	if len(users) == 0 {
		c.Data(200, "text/plain", nil)
		return
	}

	for _, user := range users {
		if userChannel := gc.remove(req.ChannelID, user); userChannel != "" {
			postMessageToSlack(req.Text, userChannel, req.UserID, "removed from group")
		}
	}

	c.JSON(200, gin.H{
		"text":          "",
		"response_type": "in_channel",
	})
}
