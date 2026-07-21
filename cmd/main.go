package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/di"
	"github.com/murashi19/koda-b8-backend1/internal/middleware"
)

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

	router.POST("/auth/register", auth.Register)
	router.POST("/auth/login", auth.Login)

	router.GET("/users", middleware.AuthMiddleware(), user.GetUsers)
	router.GET("/users/:id", middleware.AuthMiddleware(), user.GetById)
	router.POST("/users", middleware.AuthMiddleware(), user.CreateUser)
	router.PATCH("/users/:id", middleware.AuthMiddleware(), user.UpdateUser)
	router.DELETE("/users/:id", middleware.AuthMiddleware(), user.DeleteUser)

	log.Fatal(router.Run(":8080"))
}
