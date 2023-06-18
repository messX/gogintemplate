package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/messx/gogintemplate/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"live": "ok"})
	})
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	DataRoutes(route)
	UserRoutes(route)
}
