package config

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/Cingihimut/catering-apps/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	migrator := db.Migrator()

	migrator.CreateTable(&models.Users{})
	migrator.CreateTable(&models.Address{})
	migrator.CreateTable(&models.Categories{})
	migrator.CreateTable(&models.Products{})
	migrator.CreateTable(&models.ProductCategories{})
	migrator.CreateTable(&models.ProductImages{})
	migrator.CreateTable(&models.CartItem{})
	migrator.CreateTable(&models.Cart{})
	migrator.CreateTable(&models.Transactions{})

	seedDummyData(db)
	return db, nil
}

func seedDummyData(db *gorm.DB) {
	// Seed kategori
	categories := []models.Categories{
		{ID: 1, Name: "Electronics"},
		{ID: 2, Name: "Clothing"},
		// Tambahkan kategori lainnya sesuai kebutuhan
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			fmt.Println("Error seeding category:", err)
			return
		}
	}

	// Seed produk
	products := []models.Products{
		{ProductName: "Laptop", Description: "Powerful laptop", Price: 1200.00},
		{ProductName: "T-shirt", Description: "Comfortable t-shirt", Price: 20.00},
		// Tambahkan produk lainnya sesuai kebutuhan
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			fmt.Println("Error seeding product:", err)
			return
		}

		// Seed gambar untuk setiap produk
		for i := 0; i < 3; i++ {
			image := models.ProductImages{
				ID:        0,
				ProductID: product.ID,
				ImageURL:  "",
			}
			if err := db.Create(&image).Error; err != nil {
				fmt.Println("Error seeding product image:", err)
				return
			}
		}

		// Seed kategori untuk setiap produk
		for _, category := range categories {
			// Ada kemungkinan acak produk yang dimiliki oleh beberapa kategori
			if rand.Intn(2) == 0 {
				productCategory := models.ProductCategories{
					ProductID:  product.ID,
					CategoryID: category.ID,
				}
				if err := db.Create(&productCategory).Error; err != nil {
					fmt.Println("Error seeding product category:", err)
					return
				}
			}
		}
	}

	fmt.Println("Seed data successfully.")
}
