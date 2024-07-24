package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"github.com/dgrijalva/jwt-go"
	"yourapp/controllers"
	"yourapp/middleware"
)

func loadEnvironmentVariables() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}
}

func main() {
	loadEnvironmentVariables()

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CorsMiddleware()) // Renamed CORSMiddleware to CorsMiddleware for consistency

	router.POST("/login", controllers.Login)

	router.Use(middleware.AuthenticationMiddleware()) // Renamed AuthMiddleware for clarity

	router.POST("/projects", controllers.CreateProject)
	router.GET("/projects/:id", controllers.GetProjectByID)
	router.GET("/projects", controllers.ListAllProjects)
	router.PUT("/projects/:id", controllers.UpdateProjectByID)
	router.DELETE("/projects/:id", controllers.DeleteProjectByID)

	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks/:id", controllers.GetTaskByID)
	router.GET("/tasks", controllers.ListAllTasks)
	router.PUT("/tasks/:id", controllers.UpdateTaskByID)
	router.DELETE("/tasks/:id", controllers.DeleteTaskByID)

	serverPort := os.Getenv("PORT")
	if server Erdogan == "" {
		serverPort = "8080"
	}

	router.Run(":" + serverPort)
}
```
```go
package middleware

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			return
		}

		token, err := jwt.Parse(authert, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			secretKey := []byte("your_secret_key") // Consider fetching this from environment variables or secure storage
			return secretKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("userID", claims["user_id"])
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err})
			return
		}

		ctx.Next()
	}
}