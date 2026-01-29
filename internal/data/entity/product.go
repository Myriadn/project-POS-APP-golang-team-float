package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CategoryID   uint           `gorm:"not null" json:"category_id"`
	Name         string         `gorm:"size:255;not null" json:"name"`
	Description  string         `gorm:"type:text" json:"description"`
	Image        string         `gorm:"size:500" json:"image"`
	Price        float64        `gorm:"type:decimal(15,2);not null" json:"price"`
	Availability string         `gorm:"size:20;default:'in_stock'" json:"availability"` // in_stock, out_of_stock
	MenuType     string         `gorm:"size:50;default:'normal'" json:"menu_type"`      // normal, special_deals, new_year_special, desserts_and_drinks
	Stock        int            `gorm:"not null;default:0" json:"stock"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Category     Category       `gorm:"foreignKey:CategoryID" json:"category,omitzero"`
}

func (Product) TableName() string {
	return "products"
}
