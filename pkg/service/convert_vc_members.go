package service

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/data"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/proto"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type vcMembersResp struct {
	allPlayers           []*models.Player
	controllers          []*models.Player
	initiators           []*models.Player
	sentinels            []*models.Player
	duelists             []*models.Player
	excessPlayersRemoved []*proto.Player
}

func convertVCMembersToUsers(request *proto.CreateGameRequest) (*vcMembersResp, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var allPlayers []*models.Player

	var controllers []*models.Player
	var flex []*models.Player
	var sentinels []*models.Player
	var duelists []*models.Player
	var excessPlayers []*proto.Player

	vcMembers := request.VoiceChannelMembers

	for _, member := range vcMembers {
		discUser, err := data.GetPlayer(member.Id)

		if err != nil {
			log.WithError(err).Errorf("failed to get player %v from DB", member.DisplayName)
			return nil, err
		}

		if len(discUser.Roles) != len(member.Roles) {
			discUser.Roles = member.Roles
			go data.UpdateRoles(discUser)
		}

		role := InvalidRole

		for _, r := range member.Roles {
			role = getValRoleFromRoleID(r)

			if role == InvalidRole {
				continue
			}

			break
		}

		if role == InvalidRole {
			continue
		}

		switch role {
		case Initiator:
			flex = append(flex, discUser)
		case Sentinel:
			sentinels = append(sentinels, discUser)
		case Controller:
			controllers = append(controllers, discUser)
		case Duelist:
			duelists = append(duelists, discUser)
		}

		allPlayers = append(allPlayers, discUser)
	}

	if len(allPlayers) > 10 {
		playersToRemove := len(allPlayers) - 10
		var namesRemoved []string
		var possibles []*models.Player

		if len(duelists) > 4 {
			possibles = append(possibles, duelists...)
		}

		if len(flex) > 6 {
			possibles = append(possibles, flex...)
		}

		if len(sentinels) > 3 {
			possibles = append(possibles, sentinels...)
		}

		if len(controllers) > 3 {
			possibles = append(possibles, controllers...)
		}

		var toRemoveFrom []*models.Player

		if len(possibles) < playersToRemove {
			toRemoveFrom = allPlayers
		} else {
			toRemoveFrom = possibles
		}

		for i := 0; i < playersToRemove; i++ {
			index := r.Intn(len(toRemoveFrom) - 1)
			playerToRemove := toRemoveFrom[index]
			toRemoveFrom = remove(toRemoveFrom, index)
			namesRemoved = append(namesRemoved, playerToRemove.DiscordUserName)

			duelists = removeUser(duelists, playerToRemove)
			controllers = removeUser(controllers, playerToRemove)
			flex = removeUser(flex, playerToRemove)
			sentinels = removeUser(sentinels, playerToRemove)

			allPlayers = removeUser(allPlayers, playerToRemove)
			excessPlayers = append(excessPlayers, &proto.Player{
				DiscordId:   playerToRemove.DiscordID,
				DisplayName: playerToRemove.DiscordUserName,
				UserId:      playerToRemove.UserID,
				RiotId:      playerToRemove.RiotId,
				RiotTag:     playerToRemove.RiotTag,
			})
		}
	}

	return &vcMembersResp{
		allPlayers:           allPlayers,
		controllers:          controllers,
		initiators:           flex,
		sentinels:            sentinels,
		duelists:             duelists,
		excessPlayersRemoved: excessPlayers,
	}, nil
}
