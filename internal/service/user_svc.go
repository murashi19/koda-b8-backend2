package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/murashi19/koda-b8-backend1/internal/lib"
	"github.com/murashi19/koda-b8-backend1/internal/models"
	"github.com/murashi19/koda-b8-backend1/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo             *repo.UserRepo
	refreshTokenRepo *repo.RefreshTokenRepo
}

func NewUserService(repo *repo.UserRepo, refreshTokenRepo *repo.RefreshTokenRepo) *UserService {
	return &UserService{
		repo:             repo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *UserService) GetById(ctx context.Context, id int64) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("Invalid user id")
	}
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("Failed Get User By Id")
	}
	fmt.Println(data)
	return data, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, page, limit int64) ([]*models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	return s.repo.GetAllUsers(ctx, page, limit)
}

func (s *UserService) Search(ctx context.Context, keyword string, page, limit int64) ([]*models.User, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	search := strings.TrimSpace(keyword)

	if search == "" {
		return s.repo.GetAllUsers(ctx, page, limit)
	}

	return s.repo.GetUser(ctx, keyword, page, limit)
}

func (s *UserService) CreateUser(ctx context.Context, data *models.CreateUserRequest, picturePath string) (*models.User, error) {

	// Validasi input
	if data.Email == "" ||
		data.Password == "" ||
		data.Username == "" ||
		data.Phone == "" {
		return nil, errors.New("all fields are required")
	}

	if len(data.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Cek email sudah ada
	if _, err := s.repo.FindByEmail(ctx, data.Email); err == nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	// Simpan alamat gambar
	picture := picturePath

	// Buat entity User
	user := &models.User{
		Email:    data.Email,
		Password: string(hashedPassword),
		Username: data.Username,
		Phone:    data.Phone,
		Picture:  &picture,
	}

	// Simpan ke database
	return s.repo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, data *models.UpdateUserRequest) (*models.User, error) {
	_, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	if data.Email != nil {
		user, err := s.repo.FindByEmail(ctx, *data.Email)
		if err == nil && user.ID != id {
			return nil, errors.New("Email already exists")
		}
	}

	if data.Email == nil && data.Username == nil && data.Phone == nil {
		return nil, errors.New("No data to update")
	}
	if data.Email != nil && *data.Email == "" {
		return nil, errors.New("Email is Required")
	}

	if data.Username != nil && *data.Username == "" {
		return nil, errors.New("Username is required")
	}

	if data.Phone != nil && *data.Phone == "" {
		return nil, errors.New("Phone is required")
	}
	return s.repo.UpdateUser(ctx, id, data)

}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	_, err := s.repo.GetById(ctx, id)
	if err != nil {
		return errors.New("User not found")
	}
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) Upload(ctx context.Context, id int64, data *models.UpdateUserRequest) (*models.User, error) {
	user, _ := s.repo.GetById(ctx, id)

	if user == nil {
		return nil, errors.New("User not found")
	}
	return s.repo.Upload(ctx, id, *data.Picture)
}

func (s *UserService) RefreshToken(ctx context.Context, data *models.RefreshTokenRequest) (*models.RefreshTokenResponse, error) {

	// 1. Verify JWT Refresh Token
	valid, userID := lib.VerifyRefreshToken(data.RefreshToken)
	if !valid {
		return nil, errors.New("invalid refresh token")
	}

	// 2. Hash refresh token
	tokenHash := lib.HashRefreshToken(data.RefreshToken)

	// 3. Cari di database
	refreshToken, err := s.refreshTokenRepo.GetByHash(ctx, tokenHash)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// 4. Cek apakah sudah direvoke
	if refreshToken.RevokedAt != nil {
		return nil, errors.New("refresh token has been revoked")
	}

	// 5. Cek expired (double check)
	if refreshToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token has expired")
	}

	// 6. Generate access token baru
	accessToken, err := lib.GenerateAccessToken(*userID)
	if err != nil {
		return nil, err
	}

	return &models.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}
