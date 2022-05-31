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

	for index, mem := range vcMembers {
		if mem == JoviPcUserId {
			vcMembers = removeStringFromSlice(vcMembers, index)
		}
	}

	for _, member := range vcMembers {
		user, err := getMember(member)

		if err != nil {
			sendError(err.Error(), channel)
			continue
		}

		discUser := discordUser{
			userId: user.User.ID,
			nick:   user.User.Username,
			roles:  user.Roles,
		}

		role := InvalidRole
		obs := false

		for _, r := range user.Roles {
			if r == ModRoleID {
				continue
			}

			if r == ObserverRoleID {
				obs = true
				break
			}

			role = getValRoleFromRoleID(r)

			if role == InvalidRole {
				continue
			}

			break
		}

		if role == InvalidRole && !obs {
			sendError(fmt.Sprintf("Member %v, does not have a valid valorant role", user.User.Username), channel)
			continue
		}

		switch role {
		case Initiator:
			flex = append(flex, discUser)
		case Sentinel:
			sentinels = append(sentinels, discUser)
		case Controller:
			controllers = append(controllers, discUser)
		case Duelist:
			duelists = append(duelists, discUser)
		}

		if !obs {
			allPlayers = append(allPlayers, discUser)
		} else {
			sendMessage(fmt.Sprintf("%v has Observer / Caster Role, and was taken into consideration for team gen", discUser.nick), channel)
		}
	}

	if len(allPlayers) > 10 {
		playersToRemove := len(allPlayers) - 10
		msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("There are an excess number of players, randomly removing %v players", playersToRemove), channel))
		var namesRemoved []string
		var possibles []discordUser

		if len(duelists) > 4 {
			possibles = append(possibles, duelists...)
		}

		if len(flex) > 6 {
			possibles = append(possibles, flex...)
		}

		if len(sentinels) > 3 {
			possibles = append(possibles, sentinels...)
		}

		if len(controllers) > 3 {
			possibles = append(possibles, controllers...)
		}

		var toRemoveFrom []discordUser

		if len(possibles) < playersToRemove {
			//IE, not enough possible overflows
			toRemoveFrom = append(toRemoveFrom, allPlayers...)
		} else {
			toRemoveFrom = append(toRemoveFrom, possibles...)
		}

		for i := 0; i < playersToRemove; i++ {
			index := r.Intn(len(toRemoveFrom) - 1)
			playerToRemove := toRemoveFrom[index]
			toRemoveFrom = remove(toRemoveFrom, index)
			namesRemoved = append(namesRemoved, playerToRemove.nick)
		}

		logWithArgs("Names removed %v", namesRemoved)
		logWithArgs("All Players len %v", len(allPlayers))

		for _, n := range namesRemoved {
			for index, a := range allPlayers {
				if a.nick == n {
					logWithArgs("Removing player %v", n)
					allPlayers = remove(allPlayers, index)
					break
				}
			}
		}

		//Look, this isn't probably the best way of doing this, but lets be real, they are tiny objects
		//And their len will never be insane, this may take an extra 20ms or so, but I know its robust

		logWithArgs("All Players len %v", len(allPlayers))

		msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("Players NOT playing %v", namesRemoved), channel))
	}

	return allPlayers, controllers, flex, sentinels, duelists, msgIdsToRemove
}
