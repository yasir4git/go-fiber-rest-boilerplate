package config

import (
	"log"
	"os"
)

type Config struct {
	DB_HOST     string
	SECRET      string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	PORT        string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		SECRET:      os.Getenv("SECRET"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		PORT:        os.Getenv("PORT"),
	}

	if AppConfig.DB_HOST == "" {
		log.Fatal("DB_HOST environment variable is required")
	}

	if AppConfig.SECRET == "" {
		log.Fatal("SECRET environment variable is required")
	}

	if AppConfig.DB_PORT == "" {
		log.Fatal("DB_PORT environment variable is required")
	}

	if AppConfig.DB_USER == "" {
		log.Fatal("DB_USER environment variable is required")
	}

	if AppConfig.DB_PASSWORD == "" {
		log.Fatal("DB_PASSWORD environment variable is required")
	}
	if AppConfig.DB_NAME == "" {
		log.Fatal("DB_NAME environment variable is required")
	}

	if AppConfig.PORT == "" {
		log.Fatal("PORT environment variable is required")
	}
}
