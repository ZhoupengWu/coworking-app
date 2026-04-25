package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCorsPolicy(allowedOrigins string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Next()
	}
}

func EarlyExitOnPreflighRequests() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Status(http.StatusOK)

			return
		}

		ctx.Next()
	}
}
