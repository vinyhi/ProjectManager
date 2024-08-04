package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "os"
    "github.com/joho/godotenv"
)

type Task struct {
    gorm.Model
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

var db *gorm.DB
var err error

func init() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, dbUsername, dbName, dbPassword, dbPort)
    db, err = gorm.Open("postgres", dbURI)
    if err != nil {
        fmt.Println("Failed to connect to database")
    }
    db.AutoMigrate(&Task{})
}

func CreateTask(c *gin.Context) {
    var task Task
    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Create(&task)
    c.JSON(http.StatusOK, &task)
}

func GetTasks(c *gin.Context) {
    var tasks []Task
    db.Find(&tasks)

    c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
    id := c.Params.ByName("id")
    var task Task
    if err := db.Where("id = ?", id).First(&task).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
    id := c.Params.ByName("id")
    var task Task
    if err := db.Where("id = ?", id).First(&task).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Save(&task)
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Params.ByName("id")
    var task Task
    if err := db.Where("id = ?", id).Delete(&task).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully!"})
}

func main() {
    r := gin.Default()
    
    r.GET("/tasks", GetTasks)
    r.GET("/tasks/:id", GetTask)
    r.POST("/tasks", CreateTask)
    r.PUT("/tasks/:id", UpdateTask)
    r.DELETE("/tasks/:id", DeleteTask)

    r.Run()
}