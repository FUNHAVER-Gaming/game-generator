package service

import (
	"fmt"
	"testing"
)

func Test_createTeams(t *testing.T) {
	type args struct {
		controllers []discordUser
		flex        []discordUser
		sentinels   []discordUser
		duelists    []discordUser
		allPlayers  []discordUser
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Optimal Teams", args: args{
			controllers: []discordUser{
				{userId: "", nick: "zander"},
				{userId: "", nick: "zombs"}},
			flex: []discordUser{
				{userId: "", nick: "sick"},
				{userId: "", nick: "vanity"},
				{userId: "", nick: "kanpeki"},
				{userId: "", nick: "trent"},
			},
			sentinels: []discordUser{
				{userId: "", nick: "dapr"},
				{userId: "", nick: "mitch"},
			},
			duelists: []discordUser{
				{userId: "", nick: "tenz"},
				{userId: "", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "", nick: "zander"},
				{userId: "", nick: "zombs"},
				{userId: "", nick: "sick"},
				{userId: "", nick: "vanity"},
				{userId: "", nick: "kanpeki"},
				{userId: "", nick: "trent"},
				{userId: "", nick: "dapr"},
				{userId: "", nick: "mitch"},
				{userId: "", nick: "yay"},
				{userId: "", nick: "exalt"},
			},
		}},
		{name: "1 Controller but more than 4 flexes", args: args{
			controllers: []discordUser{
				{userId: "", nick: "zander"}},
			flex: []discordUser{
				{userId: "", nick: "sick"},
				{userId: "", nick: "vanity"},
				{userId: "", nick: "kanpeki"},
				{userId: "", nick: "trent"},
				{userId: "", nick: "exalt"},
			},
			sentinels: []discordUser{
				{userId: "", nick: "dapr"},
				{userId: "", nick: "mitch"},
			},
			duelists: []discordUser{
				{userId: "", nick: "tenz"},
				{userId: "", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "", nick: "zander"},
				{userId: "", nick: "sick"},
				{userId: "", nick: "vanity"},
				{userId: "", nick: "kanpeki"},
				{userId: "", nick: "trent"},
				{userId: "", nick: "dapr"},
				{userId: "", nick: "mitch"},
				{userId: "", nick: "yay"},
				{userId: "", nick: "exalt"},
				{userId: "", nick: "tenz"},
			},
		}},
		{name: "0 Controller but more than 6 flexes", args: args{
			controllers: []discordUser{},
			flex: []discordUser{
				{userId: "", nick: "sick"},
				{userId: "", nick: "vanity"},
				{userId: "", nick: "kanpeki"},
				{userId: "", nick: "trent"},
				{userId: "", nick: "exalt"},
				{userId: "", nick: "dynamic"},
			},
			sentinels: []discordUser{
				{userId: "", nick: "dapr"},
				{userId: "", nick: "mitch"},
			},
			duelists: []discordUser{
				{userId: "", nick: "tenz"},
				{userId: "", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "", nick: "zander"},
				{userId: "", nick: "sick"},
				{userId: "", nick: "vanity"},
				{userId: "", nick: "kanpeki"},
				{userId: "", nick: "trent"},
				{userId: "", nick: "dapr"},
				{userId: "", nick: "mitch"},
				{userId: "", nick: "yay"},
				{userId: "", nick: "exalt"},
				{userId: "", nick: "tenz"},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			team1, team2 := createTeams(tt.args.controllers, tt.args.flex, tt.args.sentinels, tt.args.duelists, tt.args.allPlayers)
			fmt.Printf("Team 1: %v", team1)
			fmt.Println("")
			fmt.Printf("Team 2: %v", team2)
			fmt.Println("")
		})
	}
}
