package dto

import "time"

// Auth Requests
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=4"`
}

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	OTP         string `json:"otp" binding:"required,len=4"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type CreateNewStaffManagementReq struct {
	Email             string     `json:"email" binding:"required,email"`
	Username          string     `json:"username" binding:"required,min=3,max=20"`
	Password          string     `json:"password" binding:"required,min=6"`
	FullName          string     `json:"full_name" binding:"required,min=3"`
	Phone             string     `json:"phone" binding:"required,numeric,min=10,max=15"`
	RoleID            uint       `json:"role_id" binding:"required,numeric,gt=0"`
	Salary            float64    `json:"salary" binding:"required,numeric,gte=0"`
	ShiftStart        string     `json:"shift_start" binding:"required"`
	ShiftEnd          string     `json:"shift_end" binding:"required"`
	DateOfBirth       *time.Time `json:"date_of_birth" binding:"required"`
	Address           string     `json:"address" binding:"required"`
	ProfilePicture    string     `json:"profile_picture"`
	AdditionalDetails string     `json:"additional_details"`
}

type UpdateStaffManagementReq struct {
	Email             string     `json:"email" binding:"omitempty,email"`
	Username          string     `json:"username" binding:"omitempty,min=3,max=20"`
	Password          string     `json:"password" binding:"omitempty,min=6"`
	FullName          string     `json:"full_name" binding:"omitempty,min=3"`
	Phone             string     `json:"phone" binding:"omitempty,numeric,min=10,max=15"`
	RoleID            uint       `json:"role_id" binding:"omitempty,numeric,gt=0"`
	Salary            float64    `json:"salary" binding:"omitempty,numeric,gte=0"`
	ShiftStart        string     `json:"shift_start" binding:"omitempty"`
	ShiftEnd          string     `json:"shift_end" binding:"omitempty"`
	DateOfBirth       *time.Time `json:"date_of_birth" binding:"omitempty"`
	Address           string     `json:"address"`
	ProfilePicture    string     `json:"profile_picture" binding:"omitempty"`
	AdditionalDetails string     `json:"additional_details" binding:"omitempty"`
}

type GetStaffManagementFilterRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	SortBy string `form:"sort_by"`
}
