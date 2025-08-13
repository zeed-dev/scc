package controllers

import (
	"smart-command-center-backend/services"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

// GetDashboardAnalytics returns aggregated stats for the dashboard.
func GetDashboardAnalytics(c *gin.Context) {
	stats, err := services.GetDashboardStats()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch analytics", err.Error())
		return
	}

	utils.SuccessResponse(c, "Dashboard analytics", stats)
}
