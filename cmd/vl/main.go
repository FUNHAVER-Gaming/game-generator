package main

import (
	"fmt"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/consts"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/data"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Setting up...")
	err := data.Setup(consts.PostgreConnectionString)

	if err != nil {
		log.WithError(err).Error("failed to setup database")
		return
	}

	service.Setup()
	fmt.Println("Done.")
}
