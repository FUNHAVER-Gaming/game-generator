package service

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"math/rand"
	"time"
)

func createTeams(controllers []*models.Player, initiators []*models.Player, sentinels []*models.Player, duelists []*models.Player, allPlayers []*models.Player) ([]*models.Player, []*models.Player) {
	var team1 []*models.Player
	var team2 []*models.Player

	totalDuelist := len(duelists)
	totalInitiator := len(initiators)
	totalSentinel := len(sentinels)
	totalController := len(controllers)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	//Team making logic
	if totalInitiator == OptimalInitiator && totalController == OptimalController && totalDuelist == OptimalDuelist && totalSentinel == OptimalSentinel {
		//Wow, this will never happen.
		team1 = addPlayerToTeam(team1, controllers[0])
		team2 = addPlayerToTeam(team2, controllers[1])

		team1 = addPlayerToTeam(team1, initiators[0])
		team1 = addPlayerToTeam(team1, initiators[2])
		team2 = addPlayerToTeam(team2, initiators[1])
		team2 = addPlayerToTeam(team2, initiators[3])

		team1 = addPlayerToTeam(team1, sentinels[0])
		team2 = addPlayerToTeam(team2, sentinels[1])

		team1 = addPlayerToTeam(team1, duelists[0])
		team2 = addPlayerToTeam(team2, duelists[1])
	} else {
		//I hate nested ifs, but whatever

		//Controller Block
		/*
			Priority: 2 controllers in queue, put one on each team
					  1 controller in queue, more than 4 flexes, put 1 controller and 1 initiators on each team
					  0 controller in queue, more than 6 flexes, put 1 initiators on both teams
					  0 Controller in queue, less than 6 flexes, fuck em
		*/

		//If we have an excess of controllers, lets go ahead the fill controller teams, and move them to their secondary choice
		if len(controllers) > 2 {
			randomSortAndShuffleToNew(controllers, func(role ValRole, user *models.Player) {
				removeUser(controllers, user)
				switch role {
				case Initiator:
					if !contains(initiators, user) {
						initiators = append(initiators, user)
					}
				case Duelist:
					if !contains(duelists, user) {
						duelists = append(duelists, user)
					}
				case Sentinel:
					if !contains(sentinels, user) {
						sentinels = append(sentinels, user)
					}
				}
			}, team1, team2)
		} else {
			//So, lets check if we have 2 controllers, then go from there
			if len(controllers) == 2 {
				//Ok cool, get each team with a controller
				team1 = addPlayerToTeam(team1, controllers[0])
				team2 = addPlayerToTeam(team2, controllers[1])
			}

			//OK, lets get a controller on each team but take from initiators
			if len(controllers) == 1 {
				//Wow, one controller, try to find an excess of a role
				if len(initiators) > OptimalInitiator {
					//Yay, more flexes
					index := r.Intn(len(initiators))
					flexNowController := initiators[index]
					team1 = addPlayerToTeam(team1, controllers[0])
					team2 = addPlayerToTeam(team2, flexNowController)
					remove(initiators, index)
				}
			}

			//No controllers
			if len(controllers) == 0 {
				//Do we have excess flex players?
				if len(initiators) >= 6 {
					//Assign 2 random flexes as "controllers"
					index := r.Intn(len(initiators))
					team1 = addPlayerToTeam(team1, initiators[index])
					remove(initiators, index)

					index2 := r.Intn(len(initiators))
					team2 = addPlayerToTeam(team2, initiators[index2])
					remove(initiators, index2)
				}
			}
		}

		//Now we check on duelists
		if len(duelists) > 2 {
			//Wow, more duelists than needed? _shocked_
			randomSortAndShuffleToNew(duelists, func(role ValRole, user *models.Player) {
				removeUser(duelists, user)
				switch role {
				case Initiator:
					if !contains(initiators, user) {
						initiators = append(initiators, user)
					}
				case Sentinel:
					if !contains(sentinels, user) {
						sentinels = append(sentinels, user)
					}
				}
			}, team1, team2)
		} else {
			if len(duelists) == 2 {
				//Wow, only 2 duelists in the entire lobby?
				//Ok cool, get each team with a duelist
				team1 = addPlayerToTeam(team1, duelists[0])
				team2 = addPlayerToTeam(team2, duelists[1])
			}
			//Otherwise, in this case, it doesn't matter which team they go to
		}

		//Finally, sentinels
		if len(sentinels) > 2 {
			randomSortAndShuffleToNew(sentinels, func(role ValRole, user *models.Player) {
				removeUser(sentinels, user)
				switch role {
				case Initiator:
					if !contains(initiators, user) {
						initiators = append(initiators, user)
					}
				}
			}, team1, team2)
		} else {
			if len(sentinels) == 2 {
				team1 = addPlayerToTeam(team1, sentinels[0])
				team2 = addPlayerToTeam(team2, sentinels[1])
			}
			//Otherwise, in this case, it doesn't matter which team they go to
		}
	}

	//OK, initiators. These are easier as they should've been already filtered out through everything above
	team1, team2 = randomSort(initiators, team1, team2)
	if len(team1) == 5 && len(team2) == 5 {
		//That's it, we're done
		return team1, team2
	}

	//Ok, so now teams aren't full of 5, fill from unassigned
	team1, team2 = randomSort(allPlayers, team1, team2)
	if len(team1) == 5 && len(team2) == 5 {
		//That's it, we're done
		return team1, team2
	}

	//OK, finally, for some reason, we couldn't fill all teams. Take the players left and sort them
	logWithArgs("Could NOT create properly balanced teams based off roles, despite our best efforts.")
	return randomSort(allPlayers, team1, team2)
}

func addPlayerToTeam(team []*models.Player, user *models.Player) []*models.Player {
	return append(team, user)
}
