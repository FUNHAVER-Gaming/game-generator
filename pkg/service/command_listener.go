package service

import (
	"encoding/json"
	"fmt"
	"github.com/jellydator/ttlcache/v3"
	"io/ioutil"
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

	author := request.Author
	if len(author) == 0 {
		http.Error(w, "author is empty", http.StatusBadRequest)
		return
	}

	if value := cache.Get(author); value != nil {
		if value.Value() == channel && !value.IsExpired() {
			sendError("You have sent too many requests, please wait.", channel)
			return
		}
	}

	cache.Set(author, channel, 10*time.Second)

	var msgIdsToRemove []string
	msgIdsToRemove = append(msgIdsToRemove, sendMessage("Creating game, please wait...", channel))

	defer deleteMessagesBulk(msgIdsToRemove, channel)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		voiceChannel := getVoiceChannelByTextChannel(channel)

		if len(voiceChannel) == 0 {
			sendError("No voice channel found for this text channel", channel)
			return
		}

		if err != nil {
			sendError(err.Error(), channel)
			return
		}

		if len(request.VoiceChannelMembers) < 10 {
			sendError(fmt.Sprintf("Did not find 10 users, only found %v", len(request.VoiceChannelMembers)), channel)
			return
		}

		allPlayers, controllers, flex, sentinels, duelists, msgIds := convertVCMembersToUsers(request, msgIdsToRemove, channel)
		msgIdsToRemove = append(msgIdsToRemove, msgIds...)

		msgIdsToRemove = append(msgIdsToRemove, sendMessage("Players and roles have been mapped, creating teams...", channel))
		start := time.Now().UnixNano() / int64(time.Millisecond)
		team1, team2 := createTeams(controllers, flex, sentinels, duelists, allPlayers)
		time.Sleep(1 * time.Second)
		end := time.Now().UnixNano() / int64(time.Millisecond)
		diff := end - start

		msgIdsToRemove = append(msgIdsToRemove, sendMessage(fmt.Sprintf("Created teams, took %vms. Finalizing...", diff), channel))

		printTeams(team1, team2, allPlayers, channel, err)

	}()

	wg.Wait()
	w.WriteHeader(http.StatusOK)
}
