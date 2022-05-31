package data

import "github.com/FUNHAVER-Gaming/game-generator/pkg/models"

func GetPlayerByDiscordId(discordId string) (*models.Player, error) {
	var player models.Player
	err := getAsync(&player, getPlayerFromDiscordId, discordId)

	if err != nil {
		return nil, err
	}

	cachePlayer(&player)
	return &player, nil
}
