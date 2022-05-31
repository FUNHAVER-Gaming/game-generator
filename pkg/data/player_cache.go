package data

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"github.com/jellydator/ttlcache/v3"
	"time"
)

var (
	playerCache = ttlcache.New[string, *models.Player](
		ttlcache.WithTTL[string, string](2 * time.Hour),
	)
	discordToPlayerCache = ttlcache.New[string, *models.Player](
		ttlcache.WithTTL[string, string](2 * time.Hour),
	)
)

func cachePlayer(player *models.Player) {
	playerCache.Set(player.UserID, player, 2*time.Hour)
	discordToPlayerCache.Set(player.DiscordID, player, 2*time.Hour)
}

func getCachedPlayerById(userId string) *models.Player {
	resp := playerCache.Get(userId)

	if resp == nil || resp.IsExpired() {
		return nil
	}

	return resp.Value()
}

func getCachedPlayerByDiscordId(discordId string) *models.Player {
	resp := discordToPlayerCache.Get(discordId)

	if resp == nil || resp.IsExpired() {
		return nil
	}

	return resp.Value()
}
