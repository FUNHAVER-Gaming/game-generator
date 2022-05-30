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
	var tempMaps []string
	tempMaps = append(tempMaps, maps...)
	pm := playedMaps[channel]

	for _, mp := range pm {
		for index, tm := range tempMaps {
			if mp == tm {
				tempMaps = removeStringFromSlice(tempMaps, index)
			}
		}
	}

	logWithArgs("tempMaps %v (len %v)", tempMaps, len(tempMaps))

	m := ""

	if len(tempMaps) == 1 {
		m = tempMaps[0]
		tempMaps = []string{}
		tempMaps = append(tempMaps, maps...)
		playedMaps[channel] = []string{}
	} else {
		index := r.Intn(len(tempMaps) - 1)
		m = tempMaps[index]
	}

	playMap(channel, m)
	return strings.ToTitle(m)
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
	mp = append(mp, m)
	playedMaps[channel] = mp
}
