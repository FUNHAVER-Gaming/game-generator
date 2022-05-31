package models

import "time"

type Game struct {
	GameID            string            `json:"game_id,omitempty" db:"game_id"`
	Players           []*PlayerGameData `json:"players,omitempty" db:"players"`
	Score             *string           `json:"score,omitempty" db:"score"`
	Map               string            `json:"map,omitempty" db:"map"`
	GameStartTime     *time.Time        `json:"game_start_time,omitempty" db:"game_start_time"`
	GameLength        *float64          `json:"game_length,omitempty" db:"game_length"`
	Team1Rating       uint16            `json:"team_1_rating,omitempty" db:"team_1_rating"`
	Team2Rating       uint16            `json:"team_2_rating,omitempty" db:"team_2_rating"`
	Team1StartingSide uint8             `json:"team_1_starting_side,omitempty" db:"team_1_starting_side"`
	Team2StartingSide uint8             `json:"team_2_starting_side,omitempty" db:"team_2_starting_side"`
	WinningTeam       *uint8            `json:"winning_team,omitempty" db:"winning_team"`
	TotalRoundsPlayed *uint8            `json:"total_rounds_played,omitempty" db:"total_rounds_played"`
}
