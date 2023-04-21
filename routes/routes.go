package routes

import (
	"gemm123/treetrek/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	return router
}
