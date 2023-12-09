package config

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)


type AppConfig struct{
	DB *gorm.DB
	App *gin.Engine
	Port int
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
		panic(err)
	}
}


func LoadAppConfig() *AppConfig{
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		// Default to port 8080 
		port = 8080
	}

	app := gin.Default()

	db, err := InitDB()
	if err != nil {
		panic(err)
	}

	return &AppConfig{
		Port: port,
		App: app,
		DB: db,
	}

}