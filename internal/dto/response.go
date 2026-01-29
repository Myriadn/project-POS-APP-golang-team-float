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
type DetailStaffResponse struct {
	ID             uint      `json:"id"`
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	Phone          string    `json:"phone"`
	RoleName       string    `json:"role_name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Address        string    `json:"address"`
	ProfilePicture string    `json:"profile_picture"`
	Salary         float64   `json:"salary"`
	ShiftStart     string    `json:"shift_start"`
	ShiftEnd       string    `json:"shift_end"`
}

// pagination
type Pagination struct {
	CurrentPage  int   `json:"current_page"`
	Limit        int   `json:"limit"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}

type GetlAllStaffResponse struct {
	ID       uint    `json:"id"`
	Email    string  `json:"email"`
	FullName string  `json:"full_name"`
	Phone    string  `json:"phone"`
	RoleName string  `json:"role_name"`
	Age      string  `json:"age"`
	Timing   string  `json:"timing"`
	Salary   float64 `json:"salary"`
}

// Struct Wrapper agar data & meta terbungkus rapi
type PaginationData struct {
	Items any         `json:"items"`
	Meta  *Pagination `json:"meta"`
}

// detail category menu
type DetailCategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type AllCategoryMenuResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	TotalItems int64  `json:"total_items"`
	Icon       string `json:"icon"`
}

// detail product menu
type DetailProductResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategotyName string  `json:"category_name"`
	Image        string  `json:"image"`
	Availability string  `json:"availability"`
}

// all product menu
type AllProductResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategotyName string  `json:"category_name"`
	Image        string  `json:"image"`
	Availability string  `json:"availability"`
}
