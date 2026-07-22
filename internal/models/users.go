package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Phone     string    `json:"phone"`
	Username  string    `json:"username"`
	Picture   *string   `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
	Username string `form:"username" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Username *string `json:"username"`
	Picture  *string `json:"picture"`
}
