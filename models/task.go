package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string
	Description string
	ProjectID   uint
	Assignee    string
}

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load the env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Task{})

	return db
}

func main() {
	db := ConnectDatabase()

	exampleTask := Task{Name: "Example Task", Description: "This is just an example task.", ProjectID: 1, Assignee: "John Doe"}
	db.Create(&exampleTask)
}