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

		discUser := discordUser{
			userId: user.ID,
			nick:   member.Nick,
		}

		if len(member.Roles) >= 2 {
			for _, r := range member.Roles {
				if r == ModRoleID {
					continue
				}

				valRole := getValRoleFromRoleID(r)
				if valRole == -1 {
					continue
				}

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
				continue
			}

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

		allPlayers = append(allPlayers, discUser)

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
