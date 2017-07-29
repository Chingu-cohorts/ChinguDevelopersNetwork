package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	// Required to use PostgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB returns a connection to the database
func InitDB() *gorm.DB {
	config, err := LoadSettings("config.json")
	if err != nil {
		log.Fatalf("Error loading config file for the database: %v", err)
	}

	// dbEnv would be an string like "postgres://user:pwd..."
	dbEnv := os.Getenv("DATABASE_URL")

	var dbParameters string

	if dbEnv != "" {
		dbParameters = dbEnv
	} else {
		dbParameters = "host=" + config.Database.Host +
			" user=" + config.Database.User +
			" dbname=" + config.Database.Name +
			" sslmode=disable" +
			" password=" + config.Database.Password
	}

	// Create connection to the database
	db, err := gorm.Open("postgres", dbParameters)
	if err != nil {
		log.Fatalf("Couldn't create connection to the database: %v", err)
	}

	// Use logrus as our logger
	var log = logrus.New()

	// Configure instance of logger
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel

	// Everything above info level will be logged
	log.SetLevel(log.Level)

	// Configure DB log
	db.SetLogger(log)
	db.LogMode(true)

	return db
}
