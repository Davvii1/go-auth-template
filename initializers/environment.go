package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	// Load environment variables from a .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
