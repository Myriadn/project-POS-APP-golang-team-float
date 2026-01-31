package entity

import (
	"time"
)

type Reservation struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	TableID         uint      `gorm:"not null" json:"table_id"`
	CustomerID      uint      `gorm:"not null" json:"customer_id"`
	ReservationDate time.Time `gorm:"type:date;not null" json:"reservation_date"`
	ReservationTime string    `gorm:"type:time;not null" json:"reservation_time"`
	PaxNumber       int       `gorm:"not null;default:1" json:"pax_number"`
	DepositFee      float64   `gorm:"type:decimal(15,2);default:0" json:"deposit_fee"`
	Status          string    `gorm:"size:20;default:'confirmed'" json:"status"`
	Notes           string    `gorm:"type:text" json:"notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Table           Table     `gorm:"foreignKey:TableID" json:"table,omitempty"`
	Customer        Customer  `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}

func (Reservation) TableName() string {
	return "reservations"
}

type Customer struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID string    `gorm:"size:50;unique;not null" json:"customer_id"`
	Title      string    `gorm:"size:10" json:"title"`
	FirstName  string    `gorm:"size:100;not null" json:"first_name"`
	LastName   string    `gorm:"size:100" json:"last_name"`
	Phone      string    `gorm:"size:20" json:"phone"`
	Email      string    `gorm:"size:255" json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Customer) TableName() string {
	return "customers"
}
