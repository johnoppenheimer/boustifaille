package database

import (
	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/johnoppenheimer/boustifaille/database/models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("boustifaille.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Place{})
	log.Info("Database connection initialized")
	return db.Debug()
}
