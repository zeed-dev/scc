package controllers

import (
	"smart-command-center-backend/services"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

// UpdateAmbulanceLocation handles incoming location updates from ambulances.
func UpdateAmbulanceLocation(c *gin.Context) {
	var input services.AmbulanceLocationRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	location, err := services.UpdateAmbulanceLocation(input)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update location", err.Error())
		return
	}

	utils.SuccessResponse(c, "Ambulance location updated", location)
}
