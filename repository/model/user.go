package model

import (
	"net/http"
	"time"
)

type User struct {
	Id              int    `gorm:"primaryKey" json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at,omitempty"`
	Password        string `json:"password"`
	RememberToken   string `json:"remember_token,omitempty"`
	DisplayName     string `json:"display_name,omitempty"`
	Telp            string `json:"telp,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}
type CreateUserRequest struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `gorm:"default:current_timestamp" json:"email_verified_at,omitempty"`
	Password        string    `json:"password"`
	RememberToken   string    `json:"remember_token,omitempty"`
	DisplayName     string    `json:"display_name"`
	Telp            string    `json:"telp,omitempty"`
	CreatedAt       time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
type UserResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    User   `json:"data,omitempty"`
}
type UpdateUserRequest struct {
	Id              int       `gorm:"primaryKey" json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `gorm:"default:current_timestamp" json:"email_verified_at,omitempty"`
	Password        string    `json:"password"`
	RememberToken   string    `json:"remember_token,omitempty"`
	DisplayName     string    `json:"display_name,omitempty"`
	Telp            string    `json:"telp,omitempty"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
type GetResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    []User `json:"data,omitempty"`
}
type DeleteUserRequest struct {
	Id int `json:"id"`
}
type DeletedResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   error  `json:"error,omitempty"`
}

type DataResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    []byte `json:"data,omitempty"`
	Error   error  `json:"error,omitempty"`
}

type Response func(w http.ResponseWriter)

func (u *User) TableName() string {
	return "users"
}
func (u *CreateUserRequest) TableName() string {
	return "users"
}

func (u *UpdateUserRequest) TableName() string {
	return "users"
}
