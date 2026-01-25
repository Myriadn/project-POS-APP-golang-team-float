package dto

import (
	"time"
)

// Auth Responses
type MessageResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type UserResponse struct {
	ID             uint      `json:"id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	Phone          string    `json:"phone"`
	RoleID         uint      `json:"role_id"`
	RoleName       string    `json:"role_name"`
	ProfilePicture string    `json:"profile_picture"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
}
