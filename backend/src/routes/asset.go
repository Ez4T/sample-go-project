package routes

import "github.com/gin-gonic/gin"

func AssetRoute(router *gin.Engine) {
	router.Static("/assets", "./file")
}
