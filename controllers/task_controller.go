package main

import (
    "fmt"
    "net/http"
    "os"
    "sync"
    "time"

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

var (
    db   *gorm.DB
    err  error
    taskCache struct {
        sync.Mutex
        tasks []Task
        lastUpdate time.Time
    }
)

func init() {
    // Previous init function code...
}

func CreateTask(c *gin.Context) {
    var task Task
    if err := c.BindJSON(&task); err != nil || task.Title == "" { // Added simple validation for title
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data, title is required"})
        return
    }

    if err := db.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }
    invalidateCache() // Invalidate cache on update
    c.JSON(http.StatusCreated, &task)
}

// Fetches tasks with simple caching
func GetTasks(c *gin.Context) {
    taskCache.Lock()
    defer taskCache.Unlock()

    // Check if cache is valid - Example: 30 seconds validity
    if time.Now().Sub(taskCache.lastUpdate) > 30*time.Second {
        if err := db.Find(&taskCache.tasks).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
            return 
        }
        taskCache.lastUpdate = time.Now()
    }

    c.JSON(http.StatusOK, taskCache.tasks)
}

func GetTask(c *gin.Context) {
    // Original implementation...
}

func UpdateTask(c *gin.Context) {
    id := c.Params.ByName("id")
    var task Task
    if db.Where("id = ?", id).First(&task).RecordNotFound() {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    if err := c.BindJSON(&task); err != nil || task.Title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data, title is required"})
        return
    }

    if err := db.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
        return
    }
    invalidateCache() // Invalidate cache on update
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    // Original implementation...
    invalidateCache() // Invalidate cache on update
}

func invalidateCache() {
    taskCache.Lock()
    defer taskCache.Unlock()

    taskCache.lastUpdate = time.Time{} // Invalidate cache
}

func main() {
    // Original main function code...
}