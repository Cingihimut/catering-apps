package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading env file")
	}
}
