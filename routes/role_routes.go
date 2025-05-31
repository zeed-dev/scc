package routes

import (
	"smart-command-center-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(rg *gin.RouterGroup) {
	roles := rg.Group("/roles")
	roles.GET("/", controllers.GetAllRoles)
}
