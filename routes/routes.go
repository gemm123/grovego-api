package routes

import (
	"gemm123/grovego-api/controller"
	"gemm123/grovego-api/middleware"
	"gemm123/grovego-api/repository"
	"gemm123/grovego-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepository := repository.NewRepositoyUser(db)
	userService := service.NewServiceUser(userRepository)
	userController := controller.NewControllerUser(userService)

	routeRepository := repository.NewRepositoryRoute(db)
	routeService := service.NewServiceRoute(routeRepository)
	routeController := controller.NewControllerRoute(routeService)

	api := router.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	auth.POST("/user", middleware.CheckAuthorization(), userController.User)

	route := api.Group("/route")
	route.POST("/recommendation", middleware.CheckAuthorization(), routeController.RecommendationRoute)
	route.POST("/finish", middleware.CheckAuthorization(), routeController.Finish)

	return router
}
