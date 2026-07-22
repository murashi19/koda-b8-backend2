package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/lib"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		prefix := "Bearer "
		if !strings.HasPrefix(authHeader, prefix){
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, lib.Response{
				Success: false,
				Message: "Authorization header is required",
			})
			return
		}

		token, _ := strings.CutPrefix(authHeader, prefix)
		if isValid, userID :=lib.VerifyAccessToken(token); isValid {
			ctx.Set("user_id", userID)
			ctx.Next()
			return 
		}
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
