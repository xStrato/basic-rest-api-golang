package migrations

import (
	"log"

	"github.com/xStrato/rest-api-30min/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatalln("Cannot run migrations for Book model: ", err.Error())
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln("Cannot run migrations for User model: ", err.Error())
	}
}
