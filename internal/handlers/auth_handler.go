package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/lib"
	"github.com/murashi19/koda-b8-backend1/internal/models"
	"github.com/murashi19/koda-b8-backend1/internal/service"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	username := ctx.PostForm("username")

	newUser, err := h.service.Register(ctx.Request.Context(), &models.CreateUserRequest{
		Email:    email,
		Password: password,
		Phone:    phone,
		Username: username,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, lib.Response{
		Success: true,
		Message: "Registered Success!",
		Result:  newUser,
	})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	user, err := h.service.Login(
		ctx.Request.Context(),
		&models.LoginRequest{
			Email:    email,
			Password: password,
		},
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Login Success",
		Result:  user,
	})
}

func (h *AuthHandler) GetUsers(ctx *gin.Context) {
	users, err := h.service.GetAllUsers(ctx.Request.Context())

	if err != nil {
		fmt.Println("Failed to Get Users")
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Result:  users,
	})
}
