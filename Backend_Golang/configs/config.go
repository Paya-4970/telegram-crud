package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DBDialect string
	DBDSN     string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return Config{
		AppPort:   os.Getenv("APP_PORT"),
		DBDialect: os.Getenv("DB_DIALECT"),
		DBDSN:     os.Getenv("DB_DSN"),
	}
}