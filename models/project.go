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
	ProjectRef  uint // Renamed from ProjectID for clarity that this is a reference ID to Project
}

type Project struct {
	gorm.Model
	Name        string
	Description string
	Tasks       []Task `gorm:"foreignKey:ProjectRef"` // Updated to match the renamed field in Task
}

func init() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	databaseURL := os.Getenv("DATABASE_URL") // Renamed from dsn for clarity
	database, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{}) // Renamed db to database for explicitness
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(&Project{}, &Task{}); err != nil { // Explicit error handling for AutoMigrate
		log.Fatalf("Failed to auto-migrate database schemas: %v", err)
	}
}