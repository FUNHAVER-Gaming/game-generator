package data

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	log "github.com/sirupsen/logrus"
)

func UpdateRoles(player *models.Player) {
	err := exec(updateRoles, player.Roles, player.UserID)
	if err != nil {
		log.WithError(err).Errorf("failed to update player %v (%v)", player.DiscordUserName, player.UserID)
	}
	cachePlayer(player)
}

func UpdateRating(player *models.Player) {
	err := exec(updateRating, player.Rating, player.UserID)
	if err != nil {
		log.WithError(err).Errorf("failed to update player %v (%v)", player.DiscordUserName, player.UserID)
	}
	cachePlayer(player)
}
