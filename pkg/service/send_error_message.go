package service

import (
	"fmt"
	"time"
)

func sendError(msg, channel string) {
	resp, err := botSession.ChannelMessageSend(channel, "Error: "+msg)
	if err != nil {
		fmt.Printf("error sending error message %v", err.Error())
		return
	}
	deleteMessage(channel, resp.ID, 10*time.Second)
}

func sendMessage(msg, channel string) string {
	resp, err := botSession.ChannelMessageSend(channel, msg)
	if err != nil {
		fmt.Printf("error sending error message %v", err.Error())
		return ""
	}

	return resp.ID
}

func deleteMessage(channel, id string, ttl time.Duration) {
	go func() {
		time.Sleep(ttl)
		logWithArgs("Deleting message ID %v in channel %v", id, channel)
		err := botSession.ChannelMessageDelete(channel, id)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
}
