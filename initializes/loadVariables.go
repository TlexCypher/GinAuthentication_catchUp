package initializes

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
