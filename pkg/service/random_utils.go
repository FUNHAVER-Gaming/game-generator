package service

import (
	"math/rand"
)

func randomSortAndShuffleToNew(baseList []discordUser, roleFunc func(role ValRole, user discordUser), team1 []discordUser, team2 []discordUser) {
	rand.Shuffle(len(baseList), func(i, j int) {
		baseList[i], baseList[j] = baseList[j], baseList[i]
	})

	//Assign to teams
	team1 = append(team1, baseList[0])
	team2 = append(team2, baseList[1])
	//Remove from slice
	remove(baseList, 0)
	remove(baseList, 0)

	//Assign to secondary roles
	for _, user := range baseList {
		member, _ := botSession.GuildMember(GuildID, user.userId)
		roles := member.Roles

		for _, role := range roles {
			//Remove any non valorant based roles
			if getValRoleFromRoleID(role) == -1 {
				continue
			}
			vr := getValRoleFromRoleID(role)
			roleFunc(vr, user)
		}
	}
}

func randomSort(base []discordUser, team1 []discordUser, team2 []discordUser) ([]discordUser, []discordUser) {
	rand.Shuffle(len(base), func(i, j int) {
		base[i], base[j] = base[j], base[i]
	})

	for _, user := range base {
		if contains(team1, user) {
			continue
		}

		if contains(team2, user) {
			continue
		}

		if len(team1) < len(team2) {
			team1 = addPlayerToTeam(team1, user)
		} else {
			team2 = addPlayerToTeam(team2, user)
		}

	}
	return team1, team2
}
