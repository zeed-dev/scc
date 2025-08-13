package controllers

import (
	"smart-command-center-backend/services"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get users", err.Error())
		return
	}
	utils.SuccessResponse(c, "Users retrieved successfully", users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}
	utils.SuccessResponse(c, "User retrieved successfully", user)
}

func CreateUser(c *gin.Context) {
	var input services.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	user, err := services.CreateUser(input)
	if err != nil {
		if err.Error() == "role not found" {
			utils.NotFoundResponse(c, "Role not found")
		} else {
			utils.InternalServerErrorResponse(c, "Failed to create user", err.Error())
		}
		return
	}

	utils.CreatedResponse(c, "User created successfully", user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input services.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	user, err := services.UpdateUser(id, input)
	if err != nil {
		if err.Error() == "role not found" {
			utils.NotFoundResponse(c, "Role not found")
		} else {
			utils.InternalServerErrorResponse(c, "Failed to update user", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, "User updated successfully", user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteUser(id); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete user", err.Error())
		return
	}
	utils.SuccessResponse(c, "User deleted successfully", nil)
}
