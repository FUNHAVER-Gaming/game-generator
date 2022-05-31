package data

import "github.com/FUNHAVER-Gaming/game-generator/pkg/models"

func GetPlayerByUserId(userId string) (*models.Player, error) {
	var player models.Player
	err := getAsync(&player, getPlayerFromUserId, userId)

	if err != nil {
		return nil, err
	}

	cachePlayer(&player)
	return &player, nil
}
