package routes

import (
	"gemm123/grovego-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	api := router.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", controllers.Register)

	return router
}
