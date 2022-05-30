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
		err := botSession.ChannelMessageDelete(channel, id)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
}

func deleteMessagesBulk(msgIdsToRemove []string, channel string) {
	go func() {
		time.Sleep(5 * time.Minute)
		msgIdsToRemove = append(msgIdsToRemove, sendMessage("Deleting system messages", channel))
		var removedCleaned []string

		for _, m := range msgIdsToRemove {
			if len(m) == 0 {
				continue
			}
			removedCleaned = append(removedCleaned, m)
		}

		err := botSession.ChannelMessagesBulkDelete(channel, removedCleaned)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
}
