package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func Connect() *Store {
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		panic("db_password is not in .env")
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		panic("db_password is not in .env")
	}
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		panic("db_password is not in .env")
	}

	dsn := "host=localhost user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Store{db}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
