package data

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"github.com/lib/pq"
)

// SavePlayer Insert a player into the player table
func SavePlayer(player *models.Player) error {
	cachePlayer(player)
	return execAsync(savePlayer, player.UserID, player.RiotId, player.RiotTag, player.DiscordID, player.DiscordUserName, pq.Array(player.Roles), player.Rating)
}
