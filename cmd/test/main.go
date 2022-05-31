package main

import (
	"fmt"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/data"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	file, err := os.Open("./cmd/test/dbconfig.env")
	if err != nil {
		panic(err.Error())
	}

	err = godotenv.Load(file.Name())
	if err != nil {
		panic(err.Error())
	}
	//
	//host := os.Getenv("HOST")
	//database := os.Getenv("DATABASE")
	//user := os.Getenv("USER")
	//port := os.Getenv("PORT")
	//password := os.Getenv("PASSWORD")

	err = data.Setup(os.Getenv("CONNECTION_STRING"))

	if err != nil {
		panic(err.Error())
	}

	//player, err := data.SavePlayer(&models.DiscordUser{
	//	UserID:      "123456789112345678",
	//	DisplayName: "MrFunHaver",
	//	Roles:       []string{consts.DuelistRoleId, consts.ObserverRoleId, consts.ModRoleId},
	//}, "MrFunHaver#350", 100)
	//
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println(player.UserID)
	//
	//player, err = data.GetPlayer(player.UserID)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println(player.DiscordUserName)
	//mpY4ZinM2E-VIdR1CA6Jg

	player, err := data.GetPlayer("W37Yb3oaYj-tdqRteoySB")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(player.DiscordUserName)

	return
}
