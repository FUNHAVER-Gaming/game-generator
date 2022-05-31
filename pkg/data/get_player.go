package data

import (
	"errors"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
)

func GetPlayer(id string) (*models.Player, error) {
	player := getCachedPlayerById(id)
	if player != nil {
		return player, nil
	}

	player = getCachedPlayerByDiscordId(id)
	if player != nil {
		return player, nil
	}

	player, _ = GetPlayerByUserId(id)
	if player != nil {
		return player, nil
	}

	player, _ = GetPlayerByDiscordId(id)
	if player != nil {
		return player, nil
	}

	return nil, errors.New("could not find player by that id")
}
