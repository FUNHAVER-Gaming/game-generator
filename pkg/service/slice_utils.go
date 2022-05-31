package service

import "github.com/FUNHAVER-Gaming/game-generator/pkg/models"

func remove(s []*models.Player, i int) []*models.Player {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeStringFromSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeUser(s []*models.Player, e *models.Player) []*models.Player {
	indexToRemove := -1
	for index, a := range s {
		if a.UserID == e.UserID {
			indexToRemove = index
			break
		}
	}

	if indexToRemove == -1 {
		return s
	}

	return remove(s, indexToRemove)
}

func contains(s []*models.Player, e *models.Player) bool {
	for _, a := range s {
		if a.UserID == e.UserID {
			return true
		}
	}
	return false
}
