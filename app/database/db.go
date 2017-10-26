package database

import (
	"errors"
	"time"

	"github.com/Roverr/pizza-db/app/config"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// Model defines high level operations on Pizza place's database.
type Model struct {
	conn *sqlx.DB
}

// New initalizes a new database connection
func New(cfg *config.DatabaseConfig) (*Model, error) {
	conn, err := connect(cfg)
	if err != nil {
		return nil, nil
	}
	return &Model{conn}, nil
}

func connect(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	startTime := time.Now()

	for {
		now := time.Now()
		if now.Sub(startTime) > cfg.ConnectionTimeout {
			return nil, errors.New("error: db connection timeout reached")
		}

		db, err := sqlx.Connect("mysql", cfg.URL)
		// Successful connection.
		if err == nil {
			return db, nil
		}

		log.Debugf("Sleeping for %s before retrying connection to DB", cfg.ReconnectSleepInterval)
		time.Sleep(cfg.ReconnectSleepInterval)
	}
}
