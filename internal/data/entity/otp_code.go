package entity

import (
	"time"
)

type OTPCode struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Code      string    `gorm:"size:6;not null" json:"code"`
	Type      string    `gorm:"size:20;not null" json:"type"`
	IsUsed    bool      `gorm:"default:false" json:"is_used"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
}

func (OTPCode) TableName() string {
	return "otp_codes"
}
