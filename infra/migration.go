package infra

import (
	"base-gin-go/models"
	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
