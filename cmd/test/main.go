package main

import (
	"context"
	"fmt"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/consts"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "localhost:5501", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := proto.NewLeagueServiceClient(conn)
	resp, err := client.CreatePlayer(ctx, &proto.CreatePlayerRequest{Player: &proto.Player{
		DiscordId:   "1234",
		DisplayName: "MrFunHaver",
		RiotId:      "4567",
		RiotTag:     "MrFunHaver#350",
		Roles:       []string{consts.InitiatorRoleId},
	}})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.UserId)
	return
}
