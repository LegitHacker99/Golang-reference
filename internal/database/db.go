package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *pgxpool.Pool
var GormDB *gorm.DB

func DbInit() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	DB = pool
	log.Println("Connected to database")
}

func DbClose() {
	DB.Close()
	log.Println("Disconnected from database")
}

func GormInit() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("DB Connected")

	return GormDB
}

func GormClose() {
	sqlDB, err := GormDB.DB()
	if err != nil {
		log.Fatalf("failed to get DB from GORM DB: %v", err)
	}
	sqlDB.Close()
	log.Println("DB Connection closed")
}
