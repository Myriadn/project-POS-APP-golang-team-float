package dto

import "time"

// List & detail response
type ReservationResponse struct {
	ID              uint                         `json:"id"`
	TableID         uint                         `json:"table_id"`
	CustomerID      uint                         `json:"customer_id"`
	ReservationDate string                       `json:"reservation_date"`
	ReservationTime string                       `json:"reservation_time"`
	PaxNumber       int                          `json:"pax_number"`
	DepositFee      float64                      `json:"deposit_fee"`
	Status          string                       `json:"status"`
	Notes           string                       `json:"notes"`
	CreatedAt       time.Time                    `json:"created_at"`
	UpdatedAt       time.Time                    `json:"updated_at"`
	Table           *ReservationTableResponse    `json:"table,omitempty"`
	Customer        *ReservationCustomerResponse `json:"customer,omitempty"`
}

type ReservationTableResponse struct {
	ID          uint   `json:"id"`
	TableNumber string `json:"table_number"`
	Floor       int    `json:"floor"`
	Capacity    int    `json:"capacity"`
	Status      string `json:"status"`
}

type ReservationCustomerResponse struct {
	ID         uint   `json:"id"`
	CustomerID string `json:"customer_id"`
	Title      string `json:"title"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

// Create request
// customer_id wajib, jika customer baru, frontend harus create customer dulu
// table_id wajib
// reservation_date, reservation_time, pax_number, deposit_fee, notes opsional
// status default 'confirmed'
type CreateReservationRequest struct {
	TableID         uint    `json:"table_id" binding:"required"`
	CustomerID      uint    `json:"customer_id" binding:"required"`
	ReservationDate string  `json:"reservation_date" binding:"required"` // format: YYYY-MM-DD
	ReservationTime string  `json:"reservation_time" binding:"required"` // format: HH:MM:SS
	PaxNumber       int     `json:"pax_number"`
	DepositFee      float64 `json:"deposit_fee"`
	Notes           string  `json:"notes"`
}

// Update request: hanya boleh update table_id dan status
// status: 'confirmed', 'awaited', 'cancelled'
type UpdateReservationRequest struct {
	TableID uint   `json:"table_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}
