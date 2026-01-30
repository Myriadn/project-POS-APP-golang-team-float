package usecase

import (
	"context"
	"time"

	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
)

type ReservationUsecaseInterface interface {
	ListReservations(ctx context.Context) ([]dto.ReservationResponse, error)
	GetReservationByID(ctx context.Context, id uint) (dto.ReservationResponse, error)
	CreateReservation(ctx context.Context, req dto.CreateReservationRequest) error
	UpdateReservation(ctx context.Context, id uint, req dto.UpdateReservationRequest) error
}

type ReservationUsecase struct {
	repo repository.ReservationRepository
}

func NewReservationUsecase(repo repository.ReservationRepository) ReservationUsecaseInterface {
	return &ReservationUsecase{repo: repo}
}

func (u *ReservationUsecase) ListReservations(ctx context.Context) ([]dto.ReservationResponse, error) {
	reservations, err := u.repo.ListReservations(ctx)
	if err != nil {
		return nil, err
	}
	var result []dto.ReservationResponse
	for _, r := range reservations {
		result = append(result, toReservationResponse(r))
	}
	return result, nil
}

func (u *ReservationUsecase) GetReservationByID(ctx context.Context, id uint) (dto.ReservationResponse, error) {
	r, err := u.repo.GetReservationByID(ctx, id)
	if err != nil {
		return dto.ReservationResponse{}, err
	}
	return toReservationResponse(r), nil
}

func (u *ReservationUsecase) CreateReservation(ctx context.Context, req dto.CreateReservationRequest) error {
	reservation := entity.Reservation{
		TableID:         req.TableID,
		CustomerID:      req.CustomerID,
		ReservationDate: parseDate(req.ReservationDate),
		ReservationTime: req.ReservationTime,
		PaxNumber:       req.PaxNumber,
		DepositFee:      req.DepositFee,
		Status:          "confirmed",
		Notes:           req.Notes,
	}
	return u.repo.CreateReservation(ctx, &reservation)
}

func (u *ReservationUsecase) UpdateReservation(ctx context.Context, id uint, req dto.UpdateReservationRequest) error {
	reservation, err := u.repo.GetReservationByID(ctx, id)
	if err != nil {
		return err
	}
	reservation.TableID = req.TableID
	reservation.Status = req.Status
	reservation.UpdatedAt = time.Now()
	return u.repo.UpdateReservation(ctx, &reservation)
}

func toReservationResponse(r entity.Reservation) dto.ReservationResponse {
	return dto.ReservationResponse{
		ID:              r.ID,
		TableID:         r.TableID,
		CustomerID:      r.CustomerID,
		ReservationDate: r.ReservationDate.Format("2006-01-02"),
		ReservationTime: r.ReservationTime,
		PaxNumber:       r.PaxNumber,
		DepositFee:      r.DepositFee,
		Status:          r.Status,
		Notes:           r.Notes,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
		Table: &dto.ReservationTableResponse{
			ID:          r.Table.ID,
			TableNumber: r.Table.TableNumber,
			Floor:       r.Table.Floor,
			Capacity:    r.Table.Capacity,
			Status:      r.Table.Status,
		},
		Customer: &dto.ReservationCustomerResponse{
			ID:         r.Customer.ID,
			CustomerID: r.Customer.CustomerID,
			Title:      r.Customer.Title,
			FirstName:  r.Customer.FirstName,
			LastName:   r.Customer.LastName,
			Phone:      r.Customer.Phone,
			Email:      r.Customer.Email,
		},
	}
}

func parseDate(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}
