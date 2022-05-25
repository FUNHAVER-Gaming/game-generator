package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"valorant-league/pkg/models"
)

type discordUser struct {
	userId string
	nick   string
}

func newGameHandler(w http.ResponseWriter, req *http.Request) {
	var request *models.NewGame
	decoder := json.NewDecoder(req.Body)
	decoder.UseNumber()
	err := decoder.Decode(&request)

	if err != nil {
		fmt.Println("failed to sign in user")
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			fmt.Println("failed to read all from resp body")
			return
		}

		fmt.Printf("Request Body: %v", string(body))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonRequest, _ := json.Marshal(&request)
	fmt.Println(string(jsonRequest))

	channel := request.ChannelID
	voiceChannel := getVoiceChannelByTextChannel(channel)

	if len(voiceChannel) == 0 {
		sendError("No voice channel found for this text channel", channel, botSession)
		return
	}

	currentGuild, err := botSession.Guild(GuildID)
	if err != nil {
		sendError(err.Error(), channel, botSession)
		return
	}

	if len(request.VoiceChannelMembers) != 10 {
		sendError(fmt.Sprintf("Did not find 10 users, only found %v", len(request.VoiceChannelMembers)), channel, botSession)
		return
	}

	var allPlayers []discordUser

	var controllers []discordUser
	var flex []discordUser
	var sentinels []discordUser
	var duelists []discordUser

	for _, member := range request.VoiceChannelMembers {
		user, err := botSession.User(member)

		if err != nil {
			sendError(err.Error(), channel, botSession)
			return
		}

		member, err := botSession.State.Member(currentGuild.ID, user.ID)
		role := member.Roles[0]

		if len(member.Roles) > 1 {
			for _, r := range member.Roles {
				if r == ModRoleID {
					continue
				}
				valRole := getValRoleFromRoleID(role)
				if valRole == -1 {
					continue
				}
				role = r
			}
		}

		valRole := getValRoleFromRoleID(role)

		if valRole == -1 {
			sendError(fmt.Sprintf("Member %v has role %v, but it is not a valid ValRole", member.Nick, role), channel, botSession)
			return
		}

		discUser := discordUser{
			userId: user.ID,
			nick:   member.Nick,
		}

		allPlayers = append(allPlayers, discUser)

		switch valRole {
		case Flex:
			flex = append(flex, discUser)
		case Sentinel:
			sentinels = append(sentinels, discUser)
		case Controller:
			controllers = append(controllers, discUser)
		case Duelist:
			duelists = append(duelists, discUser)
		}
	}

	team1, team2 := createTeams(controllers, flex, sentinels, duelists, allPlayers)

	team1msg := "Team 1: "
	for i, user := range team1 {
		team1msg += user.nick
		if i != len(team1)-1 {
			team1msg += ", "
		}
	}

	team2msg := "Team 2: "
	for i, user := range team2 {
		team2msg += user.nick
		if i != len(team2)-1 {
			team2msg += ", "
		}
	}

	attackingMsg := ""
	r := rand.New(rand.NewSource(time.Now().Unix()))
	if r.Intn(100) >= 50 {
		attackingMsg = "Team 1 is attack"
	} else {
		attackingMsg = "Team 2 is attack"
	}

	lobbyLeader := allPlayers[r.Intn(len(allPlayers)-1)]
	lobbyMsg := fmt.Sprintf("Lobby leader is %v", lobbyLeader.nick)

	formattedMsg := fmt.Sprintf("%v | %v\n%v\n%v", team1msg, team2msg, attackingMsg, lobbyMsg)
	_, err = botSession.ChannelMessageSend(request.ChannelID, formattedMsg)

	if err != nil {
		fmt.Printf("error sending teams message %v", err.Error())
		return
	}

	member, err := botSession.State.Member(currentGuild.ID, request.Author)
	fmt.Printf("Sent teams message for channel %v (requested by %v)", channelNameFromId(request.ChannelID), member.Nick)
	w.WriteHeader(http.StatusOK)
}

func createTeams(controllers []discordUser, flex []discordUser, sentinels []discordUser, duelists []discordUser, allPlayers []discordUser) ([]discordUser, []discordUser) {
	var team1 []discordUser
	var team2 []discordUser

	totalDuelist := len(duelists)
	totalFlex := len(flex)
	totalSentinel := len(sentinels)
	totalController := len(controllers)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	//Team making logic
	if totalFlex == OptimalFlex && totalController == OptimalController && totalDuelist == OptimalDuelist && totalSentinel == OptimalSentinel {
		//Wow, this will never happen.

		team1 = append(team1, controllers[0])
		team2 = append(team2, controllers[1])

		team1 = append(team1, flex[0])
		team1 = append(team1, flex[2])
		team2 = append(team2, flex[1])
		team2 = append(team2, flex[3])

		team1 = append(team1, sentinels[0])
		team2 = append(team2, sentinels[1])

		team1 = append(team1, duelists[0])
		team2 = append(team2, duelists[1])
	} else {
		bothTeamsHaveController := false
		//I hate nested ifs, but whatever

		//Controller Block
		/*
			Priority: 2 controllers in queue, put one on each team
					  1 controller in queue, more than 4 flexes, put 1 controller and 1 flex on each team
					  0 controller in queue, more than 6 flexes, put 1 flex on both teams
					  0 Controller in queue, less than 6 flexes, fuck em
		*/

		//So, lets check if we have 2 controllers, then go from there
		if len(controllers) == 2 {
			//Ok cool, get each team with a controller
			team1 = append(team1, controllers[0])
			team2 = append(team2, controllers[1])
			bothTeamsHaveController = true
		}

		//OK, lets get a controller on each team but take from flex
		if len(controllers) == 1 {
			//Wow, one controller, try to find an excess of a role
			if len(flex) > OptimalFlex {
				//Yay, more flexes
				index := r.Intn(len(flex))
				flexNowController := flex[index]
				team1 = append(team1, controllers[0])
				team2 = append(team2, flexNowController)
				remove(flex, index)
				bothTeamsHaveController = true
			}
		}

		if len(controllers) == 0 {
			if len(flex) >= 6 {
				index := r.Intn(len(flex))
				team1 = append(team1, flex[index])
				remove(flex, index)

				index2 := r.Intn(len(flex))
				team2 = append(team2, flex[index2])
				remove(flex, index2)

				bothTeamsHaveController = true
			}
		}

		if !bothTeamsHaveController {
			//Randomize teams from the main list
			team1, team2 = randomSort(allPlayers, team1, team2)
		} else {
			//Put flex players onto teams
			fmt.Println("Both teams have controllers, now putting rest on")
			fmt.Println(fmt.Sprintf("Team 1: %v", team1))
			fmt.Println(fmt.Sprintf("Team 2: %v", team2))
			team1, team2 = randomSort(flex, team1, team2)

			fmt.Println("Both teams have flexes, now putting duelists on")
			fmt.Println(fmt.Sprintf("Team 1: %v", team1))
			fmt.Println(fmt.Sprintf("Team 2: %v", team2))
			team1, team2 = randomSort(duelists, team1, team2)

			fmt.Println("Both teams have duelists, now putting sentinels on")
			fmt.Println(fmt.Sprintf("Team 1: %v", team1))
			fmt.Println(fmt.Sprintf("Team 2: %v", team2))
			team1, team2 = randomSort(sentinels, team1, team2)

			fmt.Println("All teams done")
		}
	}

	//Fallback, randomize all teams
	if len(team1) != len(team2) {
		fmt.Println("Randomizing all 5 teams, because no balance was found")
		rand.Shuffle(len(allPlayers), func(i, j int) {
			allPlayers[i], allPlayers[j] = allPlayers[j], allPlayers[i]
		})
		team1 = allPlayers[0:4]
		team2 = allPlayers[5:9]
	}

	return team1, team2
}

func randomSort(base []discordUser, team1 []discordUser, team2 []discordUser) ([]discordUser, []discordUser) {
	rand.Shuffle(len(base), func(i, j int) {
		base[i], base[j] = base[j], base[i]
	})

	lastTeamA := false
	for _, user := range base {
		if contains(team1, user) || contains(team2, user) {
			continue
		}
		if !lastTeamA {
			team1 = append(team1, user)
			lastTeamA = true
		} else {
			team2 = append(team2, user)
			lastTeamA = false
		}
	}
	return team1, team2
}

func remove(s []discordUser, i int) []discordUser {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func contains(s []discordUser, e discordUser) bool {
	for _, a := range s {
		if a.userId == e.userId {
			return true
		}
	}
	return false
}
