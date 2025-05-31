package routes

import (
	"smart-command-center-backend/controllers"

	"github.com/gin-gonic/gin"
)

func PanicRoutes(rg *gin.RouterGroup) {
	panic := rg.Group("/panic")
	panic.POST("", controllers.SendPanic)
}
