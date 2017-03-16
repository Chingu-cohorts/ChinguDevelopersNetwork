package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB creates a connection to the database
func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=chingucentral sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	return db
}
