package util

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// LoadEnv loads the environment variables from the .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("No .env file found, using system environment variables")
	}
	log.Log().Msg("Environment variables loaded")
}
