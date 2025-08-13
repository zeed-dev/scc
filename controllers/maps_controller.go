package controllers

import (
	"smart-command-center-backend/services"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

// GetRoute handles requests for calculating routes using Google Maps API.
func GetRoute(c *gin.Context) {
	var input services.RouteRequest
	if err := c.ShouldBindQuery(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid query parameters", err.Error())
		return
	}

	route, err := services.GetRoute(input)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get route", err.Error())
		return
	}

	utils.SuccessResponse(c, "Route fetched successfully", route)
}
