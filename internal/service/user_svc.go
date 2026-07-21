package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/murashi19/koda-b8-backend1/internal/models"
	"github.com/murashi19/koda-b8-backend1/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetAllUsers(ctx)
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

func (s *UserService) CreateUser(ctx context.Context, data *models.CreateUserRequest) (*models.User, error) {

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
