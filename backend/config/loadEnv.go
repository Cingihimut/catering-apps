package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading env file")
	}
}
