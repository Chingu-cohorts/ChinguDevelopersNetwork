package utils

import (
	"log"

	"github.com/jinzhu/gorm"
	// Required to use PostgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB returns a connection to the database
func InitDB() *gorm.DB {
	config, err := LoadSettings("config.json")
	if err != nil {
		log.Fatalf("Error loading config file for the database: %v", err)
	}

	var dbParameters = "host=" + config.Database.Host +
		" user=" + config.Database.User +
		" dbname=" + config.Database.Name +
		" sslmode=disable" +
		" password=" + config.Database.Password

	db, err := gorm.Open("postgres", dbParameters)
	if err != nil {
		log.Fatalf("Couldn't create connection to the database: %v", err)
	}

	db.LogMode(true)

	return db
}
