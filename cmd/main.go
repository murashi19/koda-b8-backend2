package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/murashi19/koda-b8-backend1/docs"
	"github.com/murashi19/koda-b8-backend1/internal/di"
	"github.com/murashi19/koda-b8-backend1/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Backend API
// @version         1.0.0
// @description     REST API for User Management built with Gin and PostgreSQL.

// @contact.name   Murashi
// @contact.email  muhamadraflishidiq@gmail.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter JWT token with **Bearer** prefix.
// Example:
// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	defer container.Close()

	auth := container.AuthHandler()
	user := container.UserHandler()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/uploads","./uploads")

	router.POST("/auth/register", auth.Register)
	router.POST("/auth/login", auth.Login)
	router.POST("/auth/refresh", user.RefreshToken)

	router.GET("/users", middleware.AuthMiddleware(), user.GetUsers)
	router.GET("/users/:id", middleware.AuthMiddleware(), user.GetById)
	router.POST("/users", middleware.AuthMiddleware(), user.CreateUser)
	router.PATCH("/users/:id", middleware.AuthMiddleware(), user.UpdateUser)
	router.DELETE("/users/:id", middleware.AuthMiddleware(), user.DeleteUser)
	router.PATCH("/users/:id/upload", middleware.AuthMiddleware(), user.Upload)
	log.Fatal(router.Run(":8080"))
}
