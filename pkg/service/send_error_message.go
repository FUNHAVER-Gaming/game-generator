package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func sendError(msg, channel string, s *discordgo.Session) {
	_, err := s.ChannelMessageSend(channel, "Error: "+msg)
	if err != nil {
		fmt.Printf("error sending error message %v", err.Error())
	}
}

func sendMessage(msg, channel string, s *discordgo.Session) {
	_, err := s.ChannelMessageSend(channel, msg)
	if err != nil {
		fmt.Printf("error sending message %v", err.Error())
	}
}
