package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/lib"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, lib.Response{
				Success: false,
				Message: "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		if token != "hello" {
			ctx.JSON(http.StatusUnauthorized, lib.Response{
				Success: false,
				Message: "Invalid authorization token",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
