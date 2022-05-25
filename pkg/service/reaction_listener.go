package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	messageId = "978889651353956393"
)

func onReact(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.MessageID != messageId {
		return
	}

	reaction := m.MessageReaction

	var role ValRole
	switch strings.ToLower(reaction.Emoji.Name) {
	case "sentinel":
		role = Sentinel
	case "initiator":
		role = Flex
	case "controller":
		role = Controller
	case "duelist":
		role = Duelist
	}

	newRoles := []string{role.getRoleId()}

	err := s.GuildMemberEdit(GuildID, m.UserID, newRoles)
	if err != nil {
		fmt.Println(err)
	}
}
