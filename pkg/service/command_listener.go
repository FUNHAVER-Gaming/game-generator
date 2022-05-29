package service

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
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

	if len(request.VoiceChannelMembers) < 10 {
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
			nick:   user.Username,
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

	team1attack := true
	r := rand.New(rand.NewSource(time.Now().Unix()))

	if r.Intn(100) >= 50 {
		team1attack = false
	}

	team1, team2 := createTeams(controllers, flex, sentinels, duelists, allPlayers)

	t1players := ""
	t2players := ""

	team1msg := "Defenders: "
	team2msg := "Defenders: "

	if team1attack {
		team1msg = "Attackers: "
	} else {
		team2msg = "Attackers: "
	}

	for i, user := range team1 {
		t1players += user.nick
		if i != len(team1)-1 {
			t1players += ", "
		}
	}

	for i, user := range team2 {
		t2players += user.nick
		if i != len(team2)-1 {
			t2players += ", "
		}
	}

	lobbyLeader := allPlayers[r.Intn(len(allPlayers)-1)]

	fields := []*discordgo.MessageEmbedField{
		{Name: team1msg, Value: t1players},
		{Name: team2msg, Value: t2players},
		{Name: "Map", Value: chooseMap(channel)},
		{Name: "Lobby Leader", Value: lobbyLeader.nick},
		{Name: "ID", Value: fmt.Sprintf("||%v||", uuid.NewString())},
	}

	embed := &discordgo.MessageEmbed{
		URL:         "http://localhost:8000",
		Type:        discordgo.EmbedTypeRich,
		Title:       "game-1: Game Created",
		Description: "New FUNHAVER Gaming Ranked Game",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       24,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "The FUNHAVER Bot",
		},
		Fields: fields,
	}

	fmt.Println(team1msg)
	fmt.Println(team2msg)

	_, err = botSession.ChannelMessageSendEmbed(channel, embed)

	if err != nil {
		fmt.Printf("error sending teams message %v", err.Error())
		return
	}

	member, err := botSession.State.Member(currentGuild.ID, request.Author)
	fmt.Printf("Sent teams message for channel %v (requested by %v)", channelNameFromId(request.ChannelID), member.Nick)
	w.WriteHeader(http.StatusOK)
}
