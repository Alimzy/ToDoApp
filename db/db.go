package db

import (
	"log"

	"gotask/config"
	"gotask/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate creates/updates the tasks table to match the struct
	if err := DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Task{})
	log.Println("Database connected and migrated.")
}
