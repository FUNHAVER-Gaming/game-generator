package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"valorant-league/pkg/health"
)

var (
	botSession                   *discordgo.Session
	voiceChannelToTextChannelMap = map[string]string{
		Game1ID: "978807561795026964",
		Game2ID: "978807838438719508",
		Game3ID: "978807822865293332",
		Game4ID: "978807716392882197",
		Game5ID: "978807859997450271",
	}
)

func Setup() {
	dg, err := discordgo.New("Bot " + DiscordToken)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	fmt.Println("Creating handlers")
	dg.AddHandler(onReact)
	dg.AddHandler(offReact)
	dg.AddHandler(addTagEd)
	fmt.Println("Handlers created, opening bot session")
	botSession = dg
	err = botSession.Open()
	defer botSession.Close()

	if err != nil {
		fmt.Println("error opening connection:", err)
		return
	}

	fmt.Println("Opened bot session, starting router")

	go func() {
		r := mux.NewRouter()
		fmt.Println("Creating path mapping")
		r.HandleFunc("/newGame", newGameHandler).Methods(http.MethodPost)
		r.HandleFunc("/health", health.Check).Methods(http.MethodGet)

		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		fmt.Println("Serving REST server")
		log.Fatal(srv.ListenAndServe())
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	fmt.Println("Gracefully shutting down")
}

func getVoiceChannelByTextChannel(channel string) string {
	return voiceChannelToTextChannelMap[channel]
}

func channelNameFromId(channelId string) string {
	switch channelId {
	case "978808082123604078":
		return "game-1"
	case "978808103942357022":
		return "game-2"
	case "978808171726508052":
		return "game-3"
	case "978808119025082398":
		return "game-4"
	case "978808224134332467":
		return "game-5"
	}
	return ""
}
