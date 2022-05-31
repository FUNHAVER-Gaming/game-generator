package data

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"strings"
)

var session *sqlx.DB

func Setup(connection string) error {
	var err error
	session, err = sqlx.Connect("postgres", connection)

	if err != nil {
		return err
	}

	log.Info("Connected to Postrge")
	log.Info("Creating tables and types")

	return createTablesAndTypes()
}

func createTablesAndTypes() error {
	if err := exec(createPlayerGameDataType); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Info("Data type exists, ignoring")
		} else {
			log.WithError(err).Error("failed to create player game data type")
			return err
		}
	}

	if err := exec(createPlayerTable); err != nil {
		log.WithError(err).Error("failed to create player table")
		return err
	}

	if err := exec(createGameTable); err != nil {
		log.WithError(err).Error("failed to create game table")
		return err
	}

	return nil
}

func exec(query string, inputs ...interface{}) error {
	_, err := session.Exec(query, inputs...)
	return err
}

func getAsync(dest interface{}, query string, inputs ...interface{}) error {
	errChan := make(chan error)
	go func(c chan error) {
		c <- session.Get(dest, query, inputs...)
	}(errChan)
	return <-errChan
}

func execAsync(query string, inputs ...interface{}) error {
	fmt.Println(fmt.Sprintf("QUERY %v, Values %v", query, inputs))
	errChan := make(chan error)

	go func(c chan error) {
		c <- exec(query, inputs...)
	}(errChan)

	return <-errChan
}
