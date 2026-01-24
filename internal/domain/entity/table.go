package entity

import (
	"time"
)

type Table struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TableNumber string    `gorm:"uniqueIndex;size:10;not null" json:"table_number"`
	Floor       int       `gorm:"not null;default:1" json:"floor"`
	Capacity    int       `gorm:"not null;default:4" json:"capacity"`
	Status      string    `gorm:"size:20;default:'available'" json:"status"` // available, occupied, reserved
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Table) TableName() string {
	return "tables"
}
