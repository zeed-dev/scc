package routes

import (
	"smart-command-center-backend/controllers"
	"smart-command-center-backend/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	 // Tambahkan ini
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // frontend
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	r.POST("/login", controllers.Login)

	authorized := r.Group("/api/v1")
	authorized.Use(middlewares.JWTAuthMiddleware())
	{
		// Role routes
		authorized.GET("/roles", controllers.GetAllRoles)

		// User routes
		authorized.GET("/users", controllers.GetUsers)
		authorized.GET("/users/:id", controllers.GetUser)
		authorized.POST("/users", controllers.CreateUser)
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)

		// Panic route
		authorized.POST("/panic", controllers.SendPanic)

	}

	return r
}
