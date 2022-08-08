package routes

import (
	"api/web/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/login", controller.GetUsers)
	router.POST("/register", controller.CreateUsers)
}
