package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func getMember(userId string) (*discordgo.Member, error) {
	member, err := botSession.State.Member(GuildID, userId)

	if err != nil {
		if err == discordgo.ErrStateNotFound {
			member, err = botSession.GuildMember(GuildID, userId)
			if err != nil {
				return nil, err
			}
			err = botSession.State.MemberAdd(member)
			if err != nil {
				return nil, err
			}
		} else {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	return member, nil
}
