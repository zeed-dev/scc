package routes

import (
	"smart-command-center-backend/controllers"

	"github.com/gin-gonic/gin"
)

// MapRoutes registers map related routes.
func MapRoutes(rg *gin.RouterGroup) {
	maps := rg.Group("/maps")
	maps.GET("/route", controllers.GetRoute)
}
