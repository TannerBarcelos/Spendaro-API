package util

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}
