package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"github.com/dgrijalva/jwt-go"
	"yourapp/controllers"
	"yourapp/middleware"
)

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}
}

func main() {
	loadEnvVariables()

	apiRouter := gin.Default()

	apiRouter.Use(gin.Logger()) 
	apiRouter.Use(gin.Recovery())
	apiRouter.Use(middleware.CORSMiddleware()) 

	apiRouter.POST("/login", controllers.Login)

	apiRouter.Use(middleware.AuthMiddleware())

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
```
```go
package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			secret := []byte("your_secret_key")
			return secret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["user_id"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err})
			return
		}

		c.Next()
	}
}