package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	Email             string         `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Username          string         `gorm:"uniqueIndex;size:100;not null" json:"username"`
	PasswordHash      string         `gorm:"size:255;not null" json:"-"`
	FullName          string         `gorm:"size:255;not null" json:"full_name"`
	Phone             string         `gorm:"size:20" json:"phone"`
	RoleID            uint           `gorm:"not null" json:"role_id"`
	ProfilePicture    string         `gorm:"size:500" json:"profile_picture"`
	Salary            float64        `gorm:"type:decimal(15,2)" json:"salary"`
	DateOfBirth       *time.Time     `json:"date_of_birth"`
	ShiftStart        string         `gorm:"size:8" json:"shift_start"`
	ShiftEnd          string         `gorm:"size:8" json:"shift_end"`
	Address           string         `gorm:"type:text" json:"address"`
	AdditionalDetails string         `gorm:"type:text" json:"additional_details"`
	IsActive          bool           `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	Role              Role           `gorm:"foreignKey:RoleID" json:"role"`
}

func (User) TableName() string {
	return "users"
}
