package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/dto"
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

// GetUsers godoc
//
//		@Summary		Get all users
//		@Description	Retrieve all users
//		@Tags			Users
//	 	@Accept			json
//		@Produce		json
//		@Security		BearerAuth
//		@Success		200	{object}	dto.UsersResponse
//		@Failure		401	{object}	dto.ErrorResponse
//		@Failure		500	{object}	dto.ErrorResponse
//		@Router			/users [get]
func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.service.GetAllUsers(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve users",
		})
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Get All Users Success",
		Result:  users,
	})
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with profile picture
//	@Tags			Users
//	@Accept			mpfd
//	@Produce		json
//	@Security		BearerAuth
//	@Param			username	formData	string	true	"Username"
//	@Param			email		formData	string	true	"Email"
//	@Param			phone		formData	string	true	"Phone number"
//	@Param			password	formData	string	true	"Password"
//	@Param			picture		formData	file	true	"Profile picture"
//	@Success		201	{object}	dto.UserDetailResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users [post]
func (h *UserHandler) CreateUser(ctx *gin.Context) {

	var req models.CreateUserRequest
	fmt.Println("Content-Type :", ctx.ContentType())
	fmt.Println("Request Header :", ctx.Request.Header.Get("Content-Type"))
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Ambil file picture
	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Picture is required",
		})
		return
	}

	// Buat folder uploads jika belum ada
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to create upload directory",
		})
		return
	}

	// Generate nama file unik
	filename := fmt.Sprintf("%d%s",
		time.Now().UnixNano(),
		filepath.Ext(file.Filename),
	)

	// Lokasi penyimpanan file
	picturePath := filepath.Join(uploadDir, filename)

	// Simpan file
	if err := ctx.SaveUploadedFile(file, picturePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to upload picture",
		})
		return
	}

	// Simpan user ke database
	newUser, err := h.service.CreateUser(
		ctx.Request.Context(),
		&req,
		picturePath,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, lib.Response{
		Success: true,
		Message: "Created new user successfully!",
		Result:  newUser,
	})
}

// GetById godoc
//
//	@Summary		Get user by ID
//	@Description	Retrieve user detail by ID
//	@Tags			Users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"User ID"
//	@Success		200	{object}	dto.UserDetailResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Router			/users/{id} [get]
func (h *UserHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	// service
	user, err := h.service.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Get User By ID Success",
		Result:  user,
	})
}

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	Update user information
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path	int							true	"User ID"
//	@Param			request	body		dto.UpdateUserRequest	true	"Update user request"
//	@Success		200		{object}	dto.UserDetailResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Router			/users/{id} [put]
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}
	user, err := h.service.UpdateUser(ctx.Request.Context(), id, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Update User Success",
		Result: gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"username":   user.Username,
			"phone":      user.Phone,
			"picture":    user.Picture,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})

}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete a user by ID
//	@Tags			Users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"User ID"
//	@Success		200	{object}	dto.SuccessResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Router			/users/{id} [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	err = h.service.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Failed to delete user",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User deleted successfully",
	})

}

// Upload godoc
//
//	@Summary		Upload user picture
//	@Description	Upload or replace a user's profile picture
//	@Tags			Users
//	@Accept			mpfd
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int		true	"User ID"
//	@Param			picture	formData	file	true	"Profile picture"
//	@Success		200		{object}	dto.UserDetailResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Router			/users/{id}/upload [post]
func (h *UserHandler) Upload(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Failed to upload picture",
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("user-picture-%d%s", id, ext)
	dst := filepath.Join("uploads", filepath.Base(filename))

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to save picture",
		})
		return
	}

	user, err := h.service.Upload(ctx.Request.Context(), id, &models.UpdateUserRequest{
		Picture: &dst,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// ← Tambahkan response sukses
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Picture uploaded successfully",
		Result:  user,
	})
}

// RefreshToken godoc
//
//	@Summary		Refresh access token
//	@Description	Generate a new access token using a valid refresh token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RefreshTokenRequest	true	"Refresh token"
//	@Success		200		{object}	dto.RefreshTokenResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Router			/auth/refresh [post]
func (h *UserHandler) RefreshToken(ctx *gin.Context) {

	var request models.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid binding JSON",
		})
		return
	}

	result, err := h.service.RefreshToken(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Message: "Invalid refresh token",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Token refreshed successfully",
		Result:  result,
	})
}
