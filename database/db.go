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
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, continuing with system environment variables")
	}

	initializeDatabase()
}

func initializeDatabase() {
	// Define PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Connect to PostgreSQL using GORM
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema and check for migration error
	if err := db.AutoMigrate(&models.Project{}, &models.Task{}); err != nil {
		log.Printf("Failed to auto-migrate database schema: %v", err)
	}
}

// GetDB exports the database connection
func GetDB() *gorm.DB {
	if db == nil {
		log.Println("Database connection has not been established.")
		// Initialize or reconnect logic can go here
		initializeDatabase() // Consider implementing a more robust check or reconnection logic
	}
	return db
}

func main() {
	// Example usage
	db := GetDB()

	if db == nil {
		log.Println("Failed to get database connection")
		return
	}

	// Work with the database, e.g., querying, inserting, updating, etc.
}