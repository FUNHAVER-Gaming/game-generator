package service

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/consts"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/data"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/proto"
	"github.com/aidarkhanov/nanoid/v2"
)

func createPlayer(request *proto.CreatePlayerRequest) (*proto.CreatePlayerResponse, error) {
	userId, _ := nanoid.New()
	player := &models.Player{
		UserID:          userId,
		RiotId:          request.Player.RiotId,
		RiotTag:         request.Player.RiotTag,
		DiscordID:       request.Player.DiscordId,
		DiscordUserName: request.Player.DisplayName,
		Roles:           request.Player.Roles,
		Rating:          uint16(consts.BaseEloRating),
	}

	err := data.SavePlayer(player)
	if err != nil {
		return nil, err
	}

	return &proto.CreatePlayerResponse{UserId: userId}, nil
}
