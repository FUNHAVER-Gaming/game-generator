package service

import (
	"errors"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/data"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/proto"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/ratings"
	"github.com/aidarkhanov/nanoid/v2"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func createGame(request *proto.CreateGameRequest) (*proto.CreateGameResponse, error) {
	if len(request.VoiceChannelMembers) < 10 {
		return &proto.CreateGameResponse{
			Ok:    false,
			Error: "Less than 10 players",
		}, errors.New("could not make a lobby, less than 10 players")
	}

	resp, err := convertVCMembersToUsers(request)

	if err != nil {
		return &proto.CreateGameResponse{
			Ok:    false,
			Error: err.Error(),
		}, err
	}

	team1, team2 := createTeams(resp.controllers, resp.initiators, resp.sentinels, resp.duelists, resp.allPlayers)

	team1StartingSide := uint8(1)
	team2StartingSide := uint8(0)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	team1attack := true

	if r.Intn(100) > 50 {
		team1attack = false
		team1StartingSide = uint8(0)
		team2StartingSide = uint8(1)
	}

	var attackers []*proto.Player
	var defenders []*proto.Player

	for _, user := range team1 {
		player := &proto.Player{
			UserId:      user.UserID,
			DiscordId:   user.DiscordID,
			DisplayName: user.DiscordUserName,
			RiotId:      user.RiotId,
			RiotTag:     user.RiotTag,
		}
		if team1attack {
			attackers = append(attackers, player)
		} else {
			defenders = append(defenders, player)
		}
	}

	for _, user := range team2 {
		player := &proto.Player{
			UserId:      user.UserID,
			DiscordId:   user.DiscordID,
			DisplayName: user.DiscordUserName,
			RiotId:      user.RiotId,
			RiotTag:     user.RiotTag,
		}
		if team1attack {
			defenders = append(defenders, player)
		} else {
			attackers = append(attackers, player)
		}
	}

	index := r.Intn(len(resp.allPlayers) - 1)
	lobbyLeader := resp.allPlayers[index]
	gameId, _ := nanoid.New()
	m := chooseMap(request.VoiceChannelId)
	gameStart := time.Now()
	team1Rating := ratings.GetTeamRating(team1)
	team2Rating := ratings.GetTeamRating(team2)

	game := &models.Game{
		GameID:            gameId,
		Map:               m,
		GameStartTime:     &gameStart,
		Team1Rating:       team1Rating,
		Team2Rating:       team2Rating,
		Team1StartingSide: team1StartingSide,
		Team2StartingSide: team2StartingSide,
	}

	err = data.CreateGame(game)
	if err != nil {
		log.WithError(err).Error("failed to create game")
		return &proto.CreateGameResponse{
			Ok:    false,
			Error: err.Error(),
		}, err
	}

	return &proto.CreateGameResponse{
		Ok:        true,
		Attackers: attackers,
		Defenders: defenders,
		Map:       m,
		LobbyLeader: &proto.Player{
			DiscordId:   lobbyLeader.DiscordID,
			DisplayName: lobbyLeader.DiscordUserName,
			UserId:      lobbyLeader.UserID,
			RiotId:      lobbyLeader.RiotId,
		},
		GameId:               gameId,
		ExcessPlayersRemoved: uint32(len(resp.excessPlayersRemoved)),
		RemovedPlayers:       resp.excessPlayersRemoved,
	}, nil
}
