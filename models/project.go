package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	ProjectRef  uint
}

type Project struct {
	gorm.Model
	Name        string
	Description string
	Tasks       []Task `gorm:"foreignKey:ProjectRef"`
}

func init() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	database, err := connectToDatabase(databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := autoMigrateDatabase(database); err != nil {
		log.Fatalf("Failed to auto-migrate database schemas: %v", err)
	}
}

func connectToDatabase(databaseURL string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
}

func autoMigrateDatabase(database *gorm.DB) error {
	return database.AutoMigrate(&Project{}, &Task{})
}