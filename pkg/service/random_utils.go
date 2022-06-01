package service

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"math/rand"
)

func randomSortAndShuffleToNew(baseList []*models.Player, roleFunc func(role ValRole, user *models.Player), team1 []*models.Player, team2 []*models.Player) {
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
	for _, member := range baseList {
		roles := member.Roles

		for _, role := range roles {
			//Remove any non valorant based roles
			if getValRoleFromRoleID(role) == InvalidRole {
				continue
			}
			vr := getValRoleFromRoleID(role)
			roleFunc(vr, member)
		}
	}
}

func randomSort(base []*models.Player, team1 []*models.Player, team2 []*models.Player) ([]*models.Player, []*models.Player) {
	rand.Shuffle(len(base), func(i, j int) {
		base[i], base[j] = base[j], base[i]
	})

	for _, user := range base {
		if contains(team1, user) || contains(team2, user) {
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

func shuffleSlice(base []*models.Player) []*models.Player {
	rand.Shuffle(len(base), func(i, j int) {
		base[i], base[j] = base[j], base[i]
	})
	return base
}
