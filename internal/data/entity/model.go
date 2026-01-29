package entity

import (
	"time"
)

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Icon        string    `gorm:"size:500" json:"icon"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "categories"
}

type Table struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TableNumber string    `gorm:"uniqueIndex;size:10;not null" json:"table_number"`
	Floor       int       `gorm:"not null;default:1" json:"floor"`
	Capacity    int       `gorm:"not null;default:4" json:"capacity"`
	Status      string    `gorm:"size:20;default:'available'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Table) TableName() string {
	return "tables"
}

type PaymentMethod struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;size:100;not null" json:"name"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}
