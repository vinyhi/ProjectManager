package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL dialect for GORM
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

// Project represents a project entity
type Project struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Task represents a task entity within a project
type Task struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ProjectID uint   `json:"project_id"`
}

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Initialize and connect to the database
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Failed to connect to database")
	}
}

func main() {
	router := gin.Default()

	// Database migration for models
	db.AutoMigrate(&Project{}, &Task{})

	setupRoutes(router)

	// Start the server
	router.Run(":8080")
}

// setupRoutes configures the router with all the required endpoints
func setupRoutes(router *gin.Engine) {
	router.GET("/projects", GetProjects)
	router.POST("/projects", CreateProject)
	router.GET("/projects/:id", GetProject)
	router.PUT("/projects/:id", UpdateProject)
	router.DELETE("/projects/:id", DeleteProject)

	router.GET("/tasks", GetTasks)
	router.POST("/tasks", CreateTask)
	router.GET("/tasks/:id", GetTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask) // fixed typo here
}

// Handlers for Projects
func GetProjects(c *gin.Context)    {}
func CreateProject(c *gin.Context)  {}
func GetProject(c *gin.Context)     {}
func UpdateProject(c *gin.Context)  {}
func DeleteProject(c *gin.Context)  {}

// Handlers for Tasks
func GetTasks(c *gin.Context)       {}
func CreateTask(c *gin.Context)     {}
func GetTask(c *gin.GroupLayout)        {}
func UpdateTask(c *gin.Context)     {}
func DeleteTask(c *gin.Context)     {} // fixed function name typo