package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configovich struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadCfg() Configovich {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env")
	}

	return Configovich{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
