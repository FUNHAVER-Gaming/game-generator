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

	member, err := s.GuildMember(GuildID, m.UserID)
	var role ValRole
	switch strings.ToLower(reaction.Emoji.Name) {
	case "sentinel":
		role = Sentinel
	case "initiator":
		role = Initiator
	case "controller":
		role = Controller
	case "duelist":
		role = Duelist
	}

	newRoles := append(member.Roles, role.getRoleId())

	err = s.GuildMemberEdit(GuildID, m.UserID, newRoles)
	if err != nil {
		fmt.Println(err)
	}
}

func offReact(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	if m.MessageID != messageId {
		return
	}

	reaction := m.MessageReaction
	member, _ := s.GuildMember(GuildID, m.UserID)

	var role ValRole
	switch strings.ToLower(reaction.Emoji.Name) {
	case "sentinel":
		role = Sentinel
	case "initiator":
		role = Initiator
	case "controller":
		role = Controller
	case "duelist":
		role = Duelist
	}

	roles := member.Roles

	for index, r := range member.Roles {
		if r == role.getRoleId() {
			roles = removeRole(roles, index)
		}
	}
	err := s.GuildMemberEdit(GuildID, m.UserID, roles)
	if err != nil {
		fmt.Println(err)
	}

}
