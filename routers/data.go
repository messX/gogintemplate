package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/controllers"
	"github.com/messx/gogintemplate/services"
)

func DataRoutes(router *gin.Engine) {
	ctrl := controllers.MainController{
		DataService: services.NewDataService(),
	}
	v1 := router.Group("/api/v1")
	v1.GET("/data", ctrl.GetAllData)
	v1.POST("/data", ctrl.Create)
}
