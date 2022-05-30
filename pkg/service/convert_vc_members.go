package service

import (
	"fmt"
	"math/rand"
	"time"
	"valorant-league/pkg/models"
)

func convertVCMembersToUsers(request *models.NewGame, msgIdsToRemove []string, channel string) ([]discordUser, []discordUser, []discordUser, []discordUser, []discordUser, []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var allPlayers []discordUser

	var controllers []discordUser
	var flex []discordUser
	var sentinels []discordUser
	var duelists []discordUser

	vcMembers := request.VoiceChannelMembers
	msgIdsToRemove = append(msgIdsToRemove, sendMessage("Getting members from VC and their roles", channel))
	msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("Found %v total VC members", len(vcMembers)), channel))

	if len(vcMembers) > 10 {
		playersToRemove := len(vcMembers) - 10
		msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("There are an excess number of players, randomly removing %v players", playersToRemove), channel))
		var namesRemoved []string

		for i := 0; i < playersToRemove; i++ {
			index := r.Intn(len(vcMembers) - 1)
			playerToRemove := vcMembers[index]
			vcMembers = removeStringFromSlice(vcMembers, index)
			member, _ := getMember(playerToRemove)
			namesRemoved = append(namesRemoved, member.User.Username)
		}

		msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("Players NOT playing %v", namesRemoved), channel))
	}

	for _, member := range vcMembers {
		user, err := botSession.User(member)

		if err != nil {
			sendError(err.Error(), channel)
			continue
		}

		member, err := getMember(user.ID)

		discUser := discordUser{
			userId: user.ID,
			nick:   user.Username,
		}

		hasValRole := false
		if len(member.Roles) >= 2 {
			for _, r := range member.Roles {
				if r == ModRoleID {
					continue
				}

				valRole := getValRoleFromRoleID(r)
				if valRole == -1 {
					continue
				}

				hasValRole = true
				switch valRole {
				case Initiator:
					flex = append(flex, discUser)
				case Sentinel:
					sentinels = append(sentinels, discUser)
				case Controller:
					controllers = append(controllers, discUser)
				case Duelist:
					duelists = append(duelists, discUser)
				}
			}
		} else if len(member.Roles) == 1 {
			valRole := getValRoleFromRoleID(member.Roles[0])
			if valRole == -1 {
				sendError(fmt.Sprintf("Member %v, does not have a valid valorant role", user.Username), channel)
				continue
			}

			hasValRole = true
			switch valRole {
			case Initiator:
				flex = append(flex, discUser)
			case Sentinel:
				sentinels = append(sentinels, discUser)
			case Controller:
				controllers = append(controllers, discUser)
			case Duelist:
				duelists = append(duelists, discUser)
			}
		}

		if !hasValRole {
			sendError(fmt.Sprintf("Member %v, does not have a valid valorant role, adding him anyway", user.Username), channel)
		}

		allPlayers = append(allPlayers, discUser)
	}
	return allPlayers, controllers, flex, sentinels, duelists, msgIdsToRemove
}
