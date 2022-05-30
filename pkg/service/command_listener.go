package service

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/jellydator/ttlcache/v3"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
	"valorant-league/pkg/models"
)

var (
	cache = ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](10 * time.Second),
	)
)

func init() {
	go cache.Start() // starts automatic expired item deletion
}

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
		fmt.Println("failed to decode NewGame request")
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			fmt.Println("failed to read all from resp body")
			return
		}

		fmt.Printf("Request Body: %v", string(body))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	channel := request.ChannelID
	if len(channel) == 0 {
		http.Error(w, "channel is empty", http.StatusBadRequest)
		return
	}

	cache.Set(channel)

	var msgIdsToRemove []string

	msgIdsToRemove = append(msgIdsToRemove, sendMessage("Creating game, please wait...", channel))

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		r := rand.New(rand.NewSource(time.Now().Unix()))
		voiceChannel := getVoiceChannelByTextChannel(channel)

		if len(voiceChannel) == 0 {
			sendError("No voice channel found for this text channel", channel)
			return
		}

		currentGuild, err := botSession.Guild(GuildID)
		if err != nil {
			sendError(err.Error(), channel)
			return
		}

		if len(request.VoiceChannelMembers) < 10 {
			sendError(fmt.Sprintf("Did not find 10 users, only found %v", len(request.VoiceChannelMembers)), channel)
			return
		}

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
				namesRemoved = append(namesRemoved, playerToRemove)
			}

			msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("Players NOT playing %v", namesRemoved), channel))
		}

		for _, member := range vcMembers {
			user, err := botSession.User(member)

			if err != nil {
				sendError(err.Error(), channel)
				continue
			}

			member, err := botSession.State.Member(currentGuild.ID, user.ID)

			if err != nil {
				if err == discordgo.ErrStateNotFound {
					member, err = botSession.GuildMember(currentGuild.ID, user.ID)
					if err != nil {
						sendError(fmt.Sprintf("Member %v, had error %v", user.Username, err.Error()), channel)
						return
					}
				} else {
					fmt.Println(err.Error())
					sendError(fmt.Sprintf("Member %v, had error %v", user.Username, err.Error()), channel)
					continue
				}
			}

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

		team1attack := true

		if r.Intn(100) >= 50 {
			team1attack = false
		}

		msgIdsToRemove = append(msgIdsToRemove, sendMessage("Players and roles have been mapped, creating teams", channel))
		start := time.Now().UnixNano() / int64(time.Millisecond)
		team1, team2 := createTeams(controllers, flex, sentinels, duelists, allPlayers)
		time.Sleep(1 * time.Second)
		end := time.Now().UnixNano() / int64(time.Millisecond)
		diff := end - start

		msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("Created teams, took %vms", diff), channel))

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
			Title:       fmt.Sprintf("Lobby %v - Game Created", channelNameFromId(channel)),
			Description: "New FUNHAVER Gaming Game",
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       24,
			Author: &discordgo.MessageEmbedAuthor{
				Name: "The FUNHAVER Bot",
			},
			Fields: fields,
		}

		_, err = botSession.ChannelMessageSendEmbed(channel, embed)

		if err != nil {
			fmt.Printf("error sending teams message %v", err.Error())
			return
		}

		go func() {
			member, _ := botSession.State.Member(currentGuild.ID, request.Author)
			fmt.Printf("Sent teams message for channel %v (requested by %v)", channelNameFromId(request.ChannelID), member.User.Username)
		}()

		go func() {
			time.Sleep(5 * time.Minute)
			msgIdsToRemove = append(msgIdsToRemove, sendMessage("Deleting system messages", channel))
			var removedCleaned []string

			for _, m := range msgIdsToRemove {
				if len(m) == 0 {
					continue
				}
				removedCleaned = append(removedCleaned, m)
			}

			err := botSession.ChannelMessagesBulkDelete(channel, removedCleaned)
			if err != nil {
				fmt.Println(err.Error())
			}
		}()
	}()

	wg.Wait()
	w.WriteHeader(http.StatusOK)
}
