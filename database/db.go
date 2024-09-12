package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"yourproject/models" // Replace with your actual project path to models
)

var db *gorm.DB

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Define PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Connect to PostgreSQL using GORM
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Project{}, &models.Task{})
}

// GetDB exports the database connection
func GetDB() *gorm.DB {
	return db
}

func main() {
	// Example usage
	db := GetDB()

	// Work with the database, e.g., querying, etc.
}