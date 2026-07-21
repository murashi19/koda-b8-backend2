package handlers

import (
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
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}

	newUser, err := h.service.Register(ctx.Request.Context(), &req)

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
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}

	user, err := h.service.Login(
		ctx.Request.Context(),
		&req,
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
		Token : "hello",
		Result:  user,
	})
}
