package routes

import (
	"smart-command-center-backend/controllers"

	"github.com/gin-gonic/gin"
)

// DashboardRoutes registers endpoints for dashboard analytics.
func DashboardRoutes(rg *gin.RouterGroup) {
	dashboard := rg.Group("/dashboard")
	dashboard.GET("/analytics", controllers.GetDashboardAnalytics)
}
