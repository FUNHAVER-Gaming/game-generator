package service

import (
	"fmt"
)

func sendError(msg, channel string) {
	_, err := botSession.ChannelMessageSend(channel, "Error: "+msg)
	if err != nil {
		fmt.Printf("error sending error message %v", err.Error())
	}
}

func sendMessage(msg, channel string) string {
	resp, err := botSession.ChannelMessageSend(channel, msg)
	if err != nil {
		fmt.Printf("error sending error message %v", err.Error())
		return ""
	}

	return resp.ID
}
