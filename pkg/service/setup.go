package service

import (
	"github.com/FUNHAVER-Gaming/game-generator/pkg/health"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	botSession *discordgo.Session
)

func Setup() {
	go func() {
		log.Info("Starting gRPC (no error next means good)")
		err := StartServer(5501)
		if err != nil {
			log.WithError(err).Error("could not start grpc server")
			return
		}
	}()

	go func() {
		log.Info("Starting mux for health check")
		r := mux.NewRouter()
		r.HandleFunc("/health", health.Check).Methods(http.MethodGet)

		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Info("Serving REST server")
		log.Fatal(srv.ListenAndServe())
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Info("Gracefully shutting down")
}
