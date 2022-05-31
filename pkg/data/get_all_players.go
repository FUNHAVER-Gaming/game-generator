package data

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
)

func GetAllPlayers() ([]*models.Player, error) {
	var dest []*models.Player
	err := session.Select(&dest, getAllPlayers)
	return dest, err
}
