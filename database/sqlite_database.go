package database

import (
	"log"
	"time"

	"github.com/xStrato/rest-api-30min/database/migrations"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var connString = "/Users/xstrato/Documents/GoLang/rest-api-30min/database/test.db"

func StartDB() {
	database, err := gorm.Open(sqlite.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	db = database
	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetConnection() *gorm.DB {
	return db
}
