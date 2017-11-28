package main

import (
	"errors"
	"math/rand"
	"time"

	"github.com/Roverr/pizza-db/app/config"
	"github.com/Roverr/pizza-db/app/database"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.New(&config.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}
	if config.DataClean {
		if err = clean(*db); err != nil {
			log.Fatal(err)
		}
		log.Infoln("Database is cleaned")
		return
	}
	conn := db.GetConnection()
	if conn == nil {
		log.Fatal(errors.New("Connection is invalid to the database"))
	}
	rand.Seed(time.Now().Unix())
	err = generate(*db)
	if err != nil {
		log.Fatal(err)
	}
	log.Infoln("Everything is generated correctly")
}
