package routes

import (
	"api/web/controller"
	"api/web/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/image-list", controller.ListAllImages)
		v1.GET("/image-list/:username", controller.ListByUserImages)
		v1.PUT("/change-name/:id", controller.ChangeImageName)
		v1.DELETE("/delete-image/:id", controller.DelImages)
		v1.POST("/upload", controller.Upload)
		v1.POST("/multi-upload", controller.MultiUpload)
	}
}
