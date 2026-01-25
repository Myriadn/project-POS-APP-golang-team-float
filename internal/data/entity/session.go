package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null;index" json:"user_id"`
	Token          uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"token"`
	IPAddress      string    `gorm:"size:45" json:"ip_address"`
	UserAgent      string    `gorm:"type:text" json:"user_agent"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	ExpiresAt      time.Time `gorm:"not null" json:"expires_at"`
	LastActivityAt time.Time `gorm:"default:now()" json:"last_activity_at"`
	CreatedAt      time.Time `json:"created_at"`
	User           User      `gorm:"foreignKey:UserID" json:"-"`
}

func (Session) TableName() string {
	return "sessions"
}
