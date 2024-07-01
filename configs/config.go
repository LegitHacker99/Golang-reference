package configs

import (
	"log"
	"os"
)

type Config struct {
	DBDriver      string
	DBSource      string
	ServerAddress string
	Port          string
}

func LoadConfig() {
	log.Print("check")
	DB_URL := os.Getenv("DATABASE_URL")
	log.Print(DB_URL)
}
