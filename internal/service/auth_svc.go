package service

import (
	"context"
	"errors"

	"github.com/murashi19/koda-b8-backend1/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Register(ctx context.Context, data *models.CreateUserRequest) (*models.User, error) {

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

	// Buat entity User
	user := &models.User{
		Email:    data.Email,
		Password: string(hashedPassword),
		Username: data.Username,
		Phone:    data.Phone,
	}

	// Simpan ke database
	return s.repo.Create(ctx, user)
}

func (s *UserService) Login(ctx context.Context, data *models.LoginRequest) (*models.User, error) {

	user, err := s.repo.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(data.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
