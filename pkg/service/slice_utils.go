package service

func remove(s []discordUser, i int) []discordUser {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeUser(s []discordUser, e discordUser) []discordUser {
	indexToRemove := -1
	for index, a := range s {
		if a.userId == e.userId && a.nick == e.nick {
			indexToRemove = index
			break
		}
	}

	if indexToRemove == -1 {
		return s
	}

	return remove(s, indexToRemove)
}

func contains(s []discordUser, e discordUser) bool {
	for _, a := range s {
		if a.userId == e.userId && a.nick == e.nick {
			return true
		}
	}
	return false
}

func removeRole(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
