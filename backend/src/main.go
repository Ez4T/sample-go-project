package main

import (
	"api/web/config"
	"api/web/middleware"
	"api/web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(middleware.CORSMiddleware())
	config.Connect()
	routes.UserRoute(router)
	routes.ApiRoute(router)
	routes.AssetRoute(router)
	router.Run("0.0.0.0:8081")
}
