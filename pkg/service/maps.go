package service

import (
	"math/rand"
	"strings"
	"time"
)

var (
	maps = []string{
		"icebox",
		"bind",
		"split",
		"fracture",
		"haven",
		"ascent",
	}
	playedMaps = map[string][]string{}
)

func chooseMap(channel string) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	m := maps[r.Intn(len(maps)-1)]
	iteration := 0
	logWithArgs("Maps len %v", len(maps))
	for ok := false; ok; ok = mapWasPlayed(channel, m) {
		logWithArgs("ok %v, iteration #%v", ok, iteration)
		m = maps[r.Intn(len(maps)-1)]
		logWithArgs("Map chosen %v", m)
		iteration++
		if iteration >= len(maps) {
			logWithArgs("Iteration equal maps, resetting")
			//All maps have been played
			playedMaps[channel] = []string{}
			break
		}
	}

	playMap(channel, m)
	return m
}

func mapWasPlayed(channel, m string) bool {
	mp := playedMaps[channel]

	if len(mp) == 0 {
		return false
	}

	for _, mps := range mp {
		if mps == strings.ToLower(m) {
			return true
		}
	}

	return false
}

func playMap(channel, m string) {
	mp := playedMaps[channel]
	logWithArgs("PLAYMAP: Maps played %v", mp)
	mp = append(mp, m)
	logWithArgs("PLAYMAP: Appended %v", mp)
	playedMaps[channel] = mp
	logWithArgs("PLAYMAP: Re-added %v", playedMaps[channel])
}
