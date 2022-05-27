package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	messageId      = "978889651353956393"
	tagEdMessageId = "979586371591213056"
	tagEdRoleId    = "979585890018013184"
)

func addTagEd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.MessageID != tagEdMessageId {
		return
	}

	member, err := s.GuildMember(GuildID, m.UserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, role := range member.Roles {
		if role == tagEdRoleId {
			return
		}
	}

	newRoles := append(member.Roles, tagEdRoleId)
	err = s.GuildMemberEdit(GuildID, m.UserID, newRoles)
	if err != nil {
		fmt.Println(err)
	}
}

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
