package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"testing"
)

func Test_createTeams(t *testing.T) {
	dg, err := discordgo.New("Bot " + DiscordToken)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	botSession = dg
	type args struct {
		controllers []discordUser
		initiators  []discordUser
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
				{userId: "276126563261612042", nick: "zander"},
				{userId: "276126563261612042", nick: "zombs"}},
			initiators: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
			},
			sentinels: []discordUser{
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
			},
			duelists: []discordUser{
				{userId: "276126563261612042", nick: "tenz"},
				{userId: "276126563261612042", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "276126563261612042", nick: "zander"},
				{userId: "276126563261612042", nick: "zombs"},
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "yay"},
				{userId: "276126563261612042", nick: "exalt"},
			},
		}},
		{name: "1 Controller but more than 4 flexes", args: args{
			controllers: []discordUser{
				{userId: "276126563261612042", nick: "zander"}},
			initiators: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "exalt"},
			},
			sentinels: []discordUser{
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
			},
			duelists: []discordUser{
				{userId: "276126563261612042", nick: "tenz"},
				{userId: "276126563261612042", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "276126563261612042", nick: "zander"},
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "yay"},
				{userId: "276126563261612042", nick: "exalt"},
				{userId: "276126563261612042", nick: "tenz"},
			},
		}},
		{name: "0 Controller but more than 6 flexes", args: args{
			controllers: []discordUser{},
			initiators: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "exalt"},
				{userId: "276126563261612042", nick: "dynamic"},
			},
			sentinels: []discordUser{
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
			},
			duelists: []discordUser{
				{userId: "276126563261612042", nick: "tenz"},
				{userId: "276126563261612042", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "276126563261612042", nick: "zander"},
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "yay"},
				{userId: "276126563261612042", nick: "exalt"},
				{userId: "276126563261612042", nick: "tenz"},
			},
		}},
		{name: "0 Controller but less than 6 flexes", args: args{
			controllers: []discordUser{},
			initiators: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dynamic"},
			},
			sentinels: []discordUser{
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "exalt"},
			},
			duelists: []discordUser{
				{userId: "276126563261612042", nick: "tenz"},
				{userId: "276126563261612042", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "yay"},
				{userId: "276126563261612042", nick: "exalt"},
				{userId: "276126563261612042", nick: "tenz"},
				{userId: "276126563261612042", nick: "dynamic"},
			},
		}},
		{name: "3 Controllers", args: args{
			controllers: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"}},
			initiators: []discordUser{
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dynamic"},
			},
			sentinels: []discordUser{
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "exalt"},
			},
			duelists: []discordUser{
				{userId: "276126563261612042", nick: "tenz"},
				{userId: "276126563261612042", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "yay"},
				{userId: "276126563261612042", nick: "exalt"},
				{userId: "276126563261612042", nick: "tenz"},
			},
		}},
		{name: "3 Controllers w/ initiator roles", args: args{
			controllers: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"}},
			initiators: []discordUser{
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dynamic"},
				{userId: "276126563261612042", nick: "kanpeki"},
			},
			sentinels: []discordUser{
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "exalt"},
			},
			duelists: []discordUser{
				{userId: "276126563261612042", nick: "tenz"},
				{userId: "276126563261612042", nick: "yay"},
			},
			allPlayers: []discordUser{
				{userId: "276126563261612042", nick: "sick"},
				{userId: "276126563261612042", nick: "vanity"},
				{userId: "276126563261612042", nick: "kanpeki"},
				{userId: "276126563261612042", nick: "trent"},
				{userId: "276126563261612042", nick: "dapr"},
				{userId: "276126563261612042", nick: "mitch"},
				{userId: "276126563261612042", nick: "yay"},
				{userId: "276126563261612042", nick: "exalt"},
				{userId: "276126563261612042", nick: "tenz"},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			team1, team2 := createTeams(tt.args.controllers, tt.args.initiators, tt.args.sentinels, tt.args.duelists, tt.args.allPlayers)

			fmt.Printf("Team 1: %v", team1)
			fmt.Println("")
			fmt.Printf("Team 2: %v", team2)
			fmt.Println("")

			if len(team1) != 5 {
				fmt.Println(fmt.Sprintf("Team 1 is %v", len(team1)))
				t.Fail()
				return
			}

			if len(team2) != 5 {
				fmt.Println(fmt.Sprintf("Team 2 is %v", len(team2)))
				t.Fail()
				return
			}

			for _, player := range team1 {
				if contains(team2, player) {
					fmt.Println(fmt.Sprintf("Player %v is on both teams", player.nick))
					t.Fail()
					return
				}
			}

		})
	}
}
