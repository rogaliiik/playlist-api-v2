package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
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
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
