package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/messx/gogintemplate/helpers"
	"github.com/messx/gogintemplate/infra/logger"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helpers.TokenValid(ctx)
		if err != nil {
			logger.Errorf("Invalid token unauthorised %v", err)
			ctx.String(http.StatusUnauthorized, "Invalid token unauthorised")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
