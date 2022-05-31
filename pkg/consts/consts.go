package consts

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	ControllerRoleId        = "1"
	InitiatorRoleId         = "2"
	DuelistRoleId           = "3"
	SentinelRoleId          = "4"
	PostgreConnectionString = ""
	BaseEloRating           = 0
)

func init() {
	file, err := os.Open("./dev.env")
	if err != nil {
		file, err = os.Open("./prod.env")
		if err != nil {
			panic("Failed to find any .env file, failing")
			return
		}
		fmt.Println("Did not find dev.env, using production server values")
	}

	if file == nil {
		panic("No .env file found, failing")
		return
	}

	err = godotenv.Load(file.Name())

	if err != nil {
		panic(fmt.Sprintf("Failed to load env file %v", err.Error()))
		return
	}

	ControllerRoleId = os.Getenv("CONTROLLER_ROLE_ID")
	InitiatorRoleId = os.Getenv("INITIATOR_ROLE_ID")
	DuelistRoleId = os.Getenv("DUELIST_ROLE_ID")
	SentinelRoleId = os.Getenv("SENTINEL_ROLE_ID")
	PostgreConnectionString = os.Getenv("POSTGRE_CONNECTION_STRING")
	BaseEloRating, err = strconv.Atoi(os.Getenv("BASE_ELO_RATING"))
	if err != nil {
		log.WithError(err).Error("failed to set base elo rating")
	}
}
