package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// Task represents a task in a project with related field.
type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);"`
	Description string `gorm:"type:text;"`
	ProjectID   uint
	Assignee    string `gorm:"type:varchar(100);"`
}

// ConnectDatabase initializes and returns a connection to the database.
func ConnectDatabase() *gorm.DB {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic("Failed to load the env file")
	}

	// Retrieve DATABASE_URL from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is not set")
	}

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Automigrate the Task struct to update schema in the database
	if err := db.AutoMigrate(&Task{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate database schema: %v", err))
	}

	return db
}

// CreateExampleTask creates an example task in the database.
func CreateExampleTask(db *gorm.DB) {
	exampleTask := Task{
		Name:        "Example Task",
		Description: "This is just an example task.",
		ProjectID:   1,
		Assignee:    "John Doe",
	}

	// Create the task and insert it into the database
	if result := db.Create(&exampleTask); result.Error != nil {
		panic(fmt.Sprintf("Failed to create example task: %v", result.Error))
	}
	fmt.Println("Example task created successfully")
}

func main() {
	// Connect to the database
	db := ConnectDatabase()
	// Create an example task
	CreateExampleTask(db)
}