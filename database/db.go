package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"yourproject/models"
)

var db *gorm.DB

func init() {
	loadDotenv()
	initializeDatabase()
}

func loadDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, continuing with system environment variables")
	}
}

func initializeDatabase() {
	dsn := buildDSN()
	connectToDatabase(dsn)
	migrateSchema()
}

func buildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
}

func connectToDatabase(dsn string) {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func migrateSchema() {
	if err := db.AutoMigrate(&models.Project{}, &models.Task{}); err != nil {
		log.Printf("Failed to auto-migrate database schema: %v", err)
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Println("Database connection has not been established. Attempting to reinitialize...")
		initializeDatabase()
	}
	return db
}

func main() {
	db := GetDB()

	if db == nil {
		log.Println("Failed to get database connection")
		return
	}
}