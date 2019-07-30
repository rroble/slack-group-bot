package main

import "fmt"

type users map[string]bool

type groupChat struct {
	Users  map[string]string
	Groups map[string]users
}

func (g *groupChat) add(groupChannel, userID string) string {
	if _, ok := g.Groups[groupChannel]; !ok {
		g.Groups[groupChannel] = make(map[string]bool, 1)
	}

	userChannel, exists := g.Users[userID]
	if !exists {
		uc, err := getUserChannel(userID)
		if err != nil {
			fmt.Printf("Cannot get user's channel: %s\n", userID)
			return ""
		}
		userChannel = uc
	}

	g.Users[userID] = userChannel
	g.Groups[groupChannel][userChannel] = true

	return userChannel
}

func (g *groupChat) remove(groupChannel, userID string) string {
	userChannel, ok := g.Users[userID]
	if !ok {
		fmt.Println("user is not in users", userID)
		return ""
	}

	if _, ok := g.Groups[groupChannel]; !ok {
		fmt.Println("group is not in groups", groupChannel)
		return ""
	}

	if _, ok := g.Groups[groupChannel][userChannel]; !ok {
		fmt.Println("userchannel is not in group channel", groupChannel)
		return ""
	}

	g.Groups[groupChannel][userChannel] = false

	return userChannel
}
