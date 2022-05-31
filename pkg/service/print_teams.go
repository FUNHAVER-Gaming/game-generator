package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func printTeams(team1 []discordUser, team2 []discordUser, allPlayers []discordUser, channel string, err error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	team1attack := true

	if r.Intn(100) >= 50 {
		team1attack = false
	}

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

	logWithArgs("PRINT TEAMS: All Players Len %v", len(allPlayers))
	index := r.Intn(len(allPlayers) - 1)
	logWithArgs("PRINT TEAMS: Index %v", index)
	lobbyLeader := allPlayers[index]

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
}
