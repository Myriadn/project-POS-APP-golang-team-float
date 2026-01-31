package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	ListReservations(ctx context.Context) ([]entity.Reservation, error)
	GetReservationByID(ctx context.Context, id uint) (entity.Reservation, error)
	CreateReservation(ctx context.Context, reservation *entity.Reservation) error
	UpdateReservation(ctx context.Context, reservation *entity.Reservation) error
}

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) ListReservations(ctx context.Context) ([]entity.Reservation, error) {
	var reservations []entity.Reservation
	err := r.db.WithContext(ctx).Preload("Table").Preload("Customer").Find(&reservations).Error
	return reservations, err
}

func (r *reservationRepository) GetReservationByID(ctx context.Context, id uint) (entity.Reservation, error) {
	var reservation entity.Reservation
	err := r.db.WithContext(ctx).Preload("Table").Preload("Customer").First(&reservation, id).Error
	return reservation, err
}

func (r *reservationRepository) CreateReservation(ctx context.Context, reservation *entity.Reservation) error {
	return r.db.WithContext(ctx).Create(reservation).Error
}

func (r *reservationRepository) UpdateReservation(ctx context.Context, reservation *entity.Reservation) error {
	return r.db.WithContext(ctx).Model(&entity.Reservation{}).Where("id = ?", reservation.ID).Updates(map[string]interface{}{
		"table_id":   reservation.TableID,
		"status":     reservation.Status,
		"updated_at": reservation.UpdatedAt,
	}).Error
}
