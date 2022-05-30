package service

import (
	"fmt"
	"strings"
	"testing"
)

func Test_chooseMap(t *testing.T) {
	tests := []struct {
		name    string
		before  func(c string)
		channel string
		want    func(s string) bool
	}{
		{name: "Map test 1", channel: "game-1", want: func(s string) bool {
			if len(s) == 0 {
				return false
			}
			for _, m := range maps {
				if m == strings.ToLower(s) {
					return true
				}
			}
			return false
		}},
		{before: func(c string) {
			for i := 0; i < len(maps); i++ {
				fmt.Println("Chose map (before): " + chooseMap(c))
			}
		}, name: "Map test, all played", channel: "game-1", want: func(s string) bool {
			if len(s) == 0 {
				return false
			}
			for _, m := range maps {
				if m == strings.ToLower(s) {
					return true
				}
			}
			return false
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(tt.channel)
			}
			got := chooseMap(tt.channel)
			if !tt.want(got) {
				t.Errorf("FAILED: chooseMap() = %v", got)
			}
			logWithArgs("Got map: %v", got)
		})
	}
}

func Test_mapWasPlayed(t *testing.T) {
	type args struct {
		channel string
		m       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapWasPlayed(tt.args.channel, tt.args.m); got != tt.want {
				t.Errorf("mapWasPlayed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playMap(t *testing.T) {
	type args struct {
		channel string
		m       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			playMap(tt.args.channel, tt.args.m)
		})
	}
}
