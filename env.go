package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EXCHANGE_RATE_APP_ID string
	ENVIRONMENT          string
}

var DefaultConfig Config

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	appId := os.Getenv("EXCHANGE_RATE_APP_ID")
	if appId == "" {
		log.Fatalln("$EXCHANGE_RATE_APP_ID must be set")
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Fatalln("$ENVIRONMENT must be set")
	}

	DefaultConfig = Config{
		EXCHANGE_RATE_APP_ID: appId,
		ENVIRONMENT:          environment,
	}

	if environment == "dev" {
		log.Println("Loaded .env")
	}
}
