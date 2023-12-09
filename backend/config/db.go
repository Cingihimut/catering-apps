package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Cingihimut/catering-apps/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Users{}, &models.Products{}, &models.Sellers{}); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}
