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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s TimeZone=Asia/Jakarta",
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

	migrator := db.Migrator()

	migrator.CreateTable(&models.Users{})
	migrator.CreateTable(&models.Address{})
	migrator.CreateTable(&models.Categories{})
	migrator.CreateTable(&models.Products{})
	migrator.CreateTable(&models.ProductImages{})
	migrator.CreateTable(&models.CartItem{})
	migrator.CreateTable(&models.Cart{})
	migrator.CreateTable(&models.Transactions{})

	return db, nil
}
