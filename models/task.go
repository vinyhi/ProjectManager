package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);"`
	Description string `gorm:"type:text;"`
	ProjectID   uint
	Assignee    string `gorm:"type:varchar(100);"`
}

func ConnectDatabase() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load the env file: ", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	log.Println("Connecting to database...")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	log.Println("Database schema migrated successfully")

	return db
}

func CreateExampleTask(db *gorm.DB) {
	exampleTask := Task{
		Name:        "Example Task",
		Description: "This is just an example task.",
		ProjectID:   1,
		Assignee:    "John Doe",
	}

	log.Println("Creating example task...")

	if result := db.Create(&exampleTask); result.Error != nil {
		log.Fatalf("Failed to create example task: %v", result.Error)
	}

	log.Println("Example task created successfully")
}

func main() {
	db := ConnectDatabase()
	CreateExampleTask(db)
}