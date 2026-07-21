package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/lib"
	"github.com/murashi19/koda-b8-backend1/internal/models"
	"github.com/murashi19/koda-b8-backend1/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.service.GetAllUsers(ctx.Request.Context())

	if err != nil {
		fmt.Println("Failed to Get Users")
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List Users",
		Result:  users,
	})
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {

	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}

	newUser, err := h.service.CreateUser(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Failed create a user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, lib.Response{
		Success: true,
		Message: "Created New user Success!",
		Result:  newUser,
	})
}

func (h *UserHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	// service
	user, err := h.service.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Message: "Not Found",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Get User By ID Success",
		Result:  user,
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	user, err := h.service.UpdateUser(ctx.Request.Context(), id, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User updated successfully",
		Result: gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"username":   user.Username,
			"phone":      user.Phone,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})

}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	err = h.service.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User deleted successfully",
	})

}
