package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/joho/godotenv"
)

type Task struct {
    gorm.Model
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

var db *gorm.DB
var err error

func init() {
    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Error loading .env file: %s\n", err)
    }

    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, dbUsername, dbName, dbPassword, dbPort)
    db, err = gorm.Open("postgres", dbURI)
    if err != nil {
        fmt.Printf("Failed to connect to database: %s\n", err)
        os.Exit(1)
    }
    if dbErr := db.AutoMigrate(&Task{}).Error; dbErr != nil {
        fmt.Printf("Failed to migrate database: %s\n", dbErr)
        os.Exit(1)
    }
}

func CreateTask(c *gin.Context) {
    var task Task
    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    if err := db.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }
    c.JSON(http.StatusCreated, &task)
}

func GetTasks(c *gin.Context) {
    var tasks []Task
    if err := db.Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
    id := c.Params.ByName("id")
    var task Task
    if err := db.Where("id = ?", id).First(&task).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
        }
        return
    }

    c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
    id := c.Params.ByName("id")
    var task Task
    if db.Where("id = ?", id).First(&task).RecordNotFound() {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    if err := db.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Params.ByName("id")
    if err := db.Where("id = ?", id).Delete(&Task{}).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func main() {
    r := gin.Default()

    r.GET("/tasks", GetTasks)
    r.GET("/tasks/:id", GetTask)
    r.POST("/tasks", CreateTask)
    r.PUT("/tasks/:id", UpdateTask)
    r.DELETE("/tasks/:id", DeleteTask)

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}