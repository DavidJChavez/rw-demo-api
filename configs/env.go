package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
