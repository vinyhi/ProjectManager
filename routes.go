package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"
	"yourapp/controllers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	router := gin.Default()

	router.POST("/projects", controllers.CreateProject)
	router.GET("/projects/:id", controllers.GetProject)
	router.GET("/projects", controllers.GetAllProjects)
	router.PUT("/projects/:id", controllers.UpdateProject)
	router.DELETE("/projects/:id", controllers.DeleteProject)

	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks/:id", controllers.GetTask)
	router.GET("/tasks", controllers.GetAllTasks)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}