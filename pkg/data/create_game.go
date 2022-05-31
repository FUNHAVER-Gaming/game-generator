package data

import "github.com/FUNHAVER-Gaming/game-generator/pkg/models"

func CreateGame(game *models.Game) error {
	return execAsync(createGame, game.GameID, game.Map, game.GameStartTime, game.Team1Rating, game.Team2Rating, game.Team1StartingSide, game.Team2StartingSide)
}
