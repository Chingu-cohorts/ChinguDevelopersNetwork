package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// Required to use PostgreSQL
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB returns a connection to the database
func InitDB() *gorm.DB {
	// config, err := LoadSettings("config.json")
	// if err != nil {
	//	log.Fatalf("Error loading config file for the database: %v", err)
	// }

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	var dbParameters = "host=" + host +
		" user=" + user +
		" dbname=" + password +
		" sslmode=disable" +
		" password=" + dbname

	db, err := gorm.Open("postgres", dbParameters)
	if err != nil {
		log.Fatalf("Couldn't create connection to the database: %v", err)
	}

	db.LogMode(true)

	return db
}
