package database

import (
	"feldrise.com/jdl/models"
	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(&models.Group{})
	database.AutoMigrate(&models.Game{})
	database.AutoMigrate(&models.GameCard{})
	database.AutoMigrate(&models.GameMode{})
}
