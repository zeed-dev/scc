package controllers

import (
	"smart-command-center-backend/services"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

func SendPanic(c *gin.Context) {
	var input services.PanicRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	userID, exist := utils.GetUserIDFromContext(c)
	if !exist {
		utils.UnauthorizedResponse(c, "User not authenticated")
		return
	}

	panicEvent, err := services.SendPanic(userID, input)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to send panic", err.Error())
		return
	}

	utils.CreatedResponse(c, "Panic event created successfully", panicEvent)
}
