package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"
	"yourapp/controllers"
)

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}
}

func main() {
	loadEnvVariables()

	apiRouter := gin.Default()

	apiRouter.POST("/projects", controllers.CreateProject)
	apiRouter.GET("/projects/:id", controllers.GetProjectByID)
	apiRouter.GET("/projects", controllers.ListAllProjects)
	apiRouter.PUT("/projects/:id", controllers.UpdateProjectByID)
	apiRouter.DELETE("/projects/:id", controllers.DeleteProjectByID)

	apiRouter.POST("/tasks", controllers.CreateTask)
	apiRouter.GET("/tasks/:id", controllers.GetTaskByID)
	apiRouter.GET("/tasks", controllers.ListAllTasks)
	apiRouter.PUT("/tasks/:id", controllers.UpdateTaskByID)
	apiRouter.DELETE("/tasks/:id", controllers.DeleteTaskByID)

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	apiRouter.Run(":" + serverPort)
}