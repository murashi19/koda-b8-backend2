package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/dto"
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


// Register godoc
// 
// @Summary Register New User
// @Description Create a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "Register Request"
// @Success 200 {object} dto.UsersResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req models.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}

	newUser, err := h.service.Register(ctx.Request.Context(), &req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Failed to register user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, lib.Response{
		Success: true,
		Message: "Registered Success!",
		Result:  newUser,
	})
}



// Login godoc
// 
// @Summary Login User
// @Description Authenticate user using email and password, then return JWT token.
// @Tags Authentication
// @Accept json
// @Produce json 
// @Param request body models.LoginRequest true "Login Request"
// @Success 200 {object} dto.LoginResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}

	loginResponse, err := h.service.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Login success",
		Result:  loginResponse,
	})
}
