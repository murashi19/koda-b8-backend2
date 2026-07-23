package dto

import "time"

type UserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDetailResponse struct {
	Success bool         `json:"success" example:"true"`
	Message string       `json:"message" example:"Get user success"`
	Result  UserResponse `json:"result"`
}

type UsersResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Get all users success"`
	Result  []UserResponse `json:"result"`
}

type SuccessResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Unauthorized"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	Success bool          `json:"success" example:"true"`
	Message string        `json:"message" example:"Login success"`
	Result  TokenResponse `json:"result"`
}

type RefreshTokenResponse struct {
	Success bool          `json:"success" example:"true"`
	Message string        `json:"message" example:"Token refreshed successfully"`
	Result  TokenResponse `json:"result"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type UserQuery struct {
	Page   int
	Limit  int
	Search map[string]string
}
