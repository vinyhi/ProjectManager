package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/joho/godotenv"
    "os"
)

var db *gorm.DB
var err error

type Project struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

type Task struct {
    ID        uint   `json:"id"`
    Name      string `json:"name"`
    ProjectID uint   `json:"project_id"`
}

func init() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        panic("Failed to connect to database")
    }
}

func main() {
    router := gin.Default()

    db.AutoMigrate(&Project{}, &Task{})

    router.GET("/projects", GetProjects)
    router.POST("/projects", CreateProject)
    router.GET("/projects/:id", GetProject)
    router.PUT("/projects/:id", UpdateProject)
    router.DELETE("/projects/:id", DeleteProject)

    router.GET("/tasks", GetTasks)
    router.POST("/tasks", CreateTask)
    router.GET("/tasks/:id", GetTask)
    router.PUT("/tasks/:id", UpdateTask)
    router.DELETE("/tasks/:i", DeleteTak)

    router.Run(":8080")
}

func GetProjects(c *gin.Context) {}
func CreateProject(c *gin.Context) {}
func GetProject(c *gin.Context) {}
func UpdateProject(c *gin.Context) {}
func DeleteProject(c *gin.Context) {}
func GetTasks(c *gin.Context) {}
func CreateTask(c *gin.Context) {}
func GetTask(c *gin.Context) {}
func UpdateTask(c *gin.Context) {}
func DeleteTak(c *gin.Context) {}