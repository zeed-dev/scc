package routes

import (
	"smart-command-center-backend/controllers"

	"github.com/gin-gonic/gin"
)

// AmbulanceRoutes registers routes for ambulance tracking.
func AmbulanceRoutes(rg *gin.RouterGroup) {
	ambulance := rg.Group("/ambulance")
	ambulance.POST("/track", controllers.UpdateAmbulanceLocation)
}
