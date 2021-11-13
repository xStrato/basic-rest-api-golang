package migrations

import (
	"log"

	"github.com/xStrato/rest-api-30min/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatalln("Cannot run migrations: ", err.Error())
	}
}
