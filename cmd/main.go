package main

import (
	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/di"
)

func main() {
	r := gin.Default()

	c := di.NewContainer()
	authHandler := c.AuthHandler()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.Run("0.0.0.0:8080")
}
