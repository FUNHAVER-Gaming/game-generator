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

	member, err := botSession.GuildMember(GuildID, m.UserID)
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

	for _, role := range member.Roles {
		if role == ModRoleID {
			newRoles = append(newRoles, ModRoleID)
		}
	}

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
	member, _ := botSession.GuildMember(GuildID, m.UserID)

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

func removeRole(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
