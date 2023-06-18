package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/controllers"
	"github.com/messx/gogintemplate/middlewares"
	"github.com/messx/gogintemplate/services"
)

func UserRoutes(router *gin.Engine) {
	ctrl := controllers.UserController{
		UserService: services.NewUserService(),
	}
	v1 := router.Group("/api/v1")

	v1.POST("/user/register", ctrl.Register)
	v1.POST("/user/login", ctrl.Login)
	protected := router.Group("/api/v1/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user/info", ctrl.GetUser)
}
