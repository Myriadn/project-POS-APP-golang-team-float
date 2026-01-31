package entity

import (
	"time"
)

type Notification struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	Title     string     `gorm:"size:255;not null" json:"title"`
	Message   string     `gorm:"type:text" json:"message"`
	Type      string     `gorm:"size:50;default:'info'" json:"type"`
	Status    string     `gorm:"size:20;default:'new'" json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	ReadAt    *time.Time `json:"read_at"`
	User      User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}
