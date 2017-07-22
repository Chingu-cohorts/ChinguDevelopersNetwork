package utils

import (
	"log"

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

	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")

	var dbParameters = "postgres://deoggavbdbyvbz:69aa072ec89331d558d935ae85488983b482975701a7a01c9eaccd9ecfd2c524@ec2-54-243-107-66.compute-1.amazonaws.com:5432/de2rrpjmckuilt"

	db, err := gorm.Open("postgres", dbParameters)
	if err != nil {
		log.Fatalf("Couldn't create connection to the database: %v", err)
	}

	db.LogMode(true)

	return db
}
