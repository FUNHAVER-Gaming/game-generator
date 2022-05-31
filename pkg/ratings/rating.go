package ratings

import "github.com/FUNHAVER-Gaming/game-generator/pkg/models"

func GetTeamRating(team []*models.Player) uint16 {
	rating := uint16(0)

	for _, p := range team {
		rating += p.Rating
	}

	avgRating := int(rating) / len(team)
	return uint16(avgRating)
}
