package models

import (
	"github.com/lib/pq"
	"time"
)

type Player struct {
	UserID                  string         `json:"user_id,omitempty" db:"user_id"`
	RiotId                  string         `json:"riot_id,omitempty" db:"riot_id"`
	RiotTag                 string         `json:"riot_tag" db:"riot_tag"`
	DiscordID               string         `json:"discord_id,omitempty" db:"discord_id"`
	DiscordUserName         string         `json:"discord_user_name,omitempty" db:"discord_username"`
	Roles                   pq.StringArray `json:"roles,omitempty" db:"roles"`
	Rating                  uint16         `json:"rating,omitempty" db:"rating"`
	TotalGamesPlayed        *int           `json:"total_games_played,omitempty" db:"total_games_played"`
	LastGamePlayedTimestamp *time.Time     `json:"last_game_played_timestamp,omitempty" db:"last_game_played_timestamp"`
	LastGamePlayedID        *string        `json:"last_game_played_id,omitempty" db:"last_game_played_id"`
	WonRounds               *int           `json:"won_rounds,omitempty" db:"won_rounds"`
	LostRounds              *int           `json:"lost_rounds,omitempty" db:"lost_rounds"`
	TotalRounds             *int           `json:"total_rounds,omitempty" db:"total_rounds"`
	AllGamesPlayed          pq.StringArray `json:"all_games_played,omitempty" db:"all_games_played"`
}

type PlayerGameData struct {
	UserID      string `json:"user_id,omitempty" db:"user_id"`
	Agent       string `json:"agent,omitempty" db:"agent"`
	ACS         uint16 `json:"acs,omitempty" db:"acs"`
	Kills       uint16 `json:"kills,omitempty" db:"kills"`
	Deaths      uint16 `json:"deaths,omitempty" db:"deaths"`
	Assists     uint16 `json:"assists,omitempty" db:"assists"`
	FirstBloods uint16 `json:"first_bloods,omitempty" db:"first_bloods"`
	FirstDeaths uint16 `json:"first_deaths,omitempty" db:"first_deaths"`
	Plants      uint16 `json:"plants,omitempty" db:"plants"`
	Defuses     uint16 `json:"defuses,omitempty" db:"defuses"`
}
