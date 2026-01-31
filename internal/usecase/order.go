package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
)

type OrderUsecaseInterface interface {
	ListOrders(ctx context.Context) ([]dto.OrderResponse, error)
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID uint) (dto.OrderResponse, error)
	UpdateOrder(ctx context.Context, id int, req dto.UpdateOrderRequest) (dto.OrderResponse, error)
	DeleteOrder(ctx context.Context, id int) error
	ListAvailableTables(ctx context.Context) ([]dto.TableResponse, error)
	ListPaymentMethods(ctx context.Context) ([]dto.PaymentMethodResponse, error)
}

type OrderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) OrderUsecaseInterface {
	return &OrderUsecase{repo: repo}
}

func (o *OrderUsecase) ListOrders(ctx context.Context) ([]dto.OrderResponse, error) {
	return o.repo.ListOrders(ctx)
}

func (o *OrderUsecase) CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID uint) (dto.OrderResponse, error) {
	return o.repo.CreateOrder(ctx, req, userID)
}

func (o *OrderUsecase) UpdateOrder(ctx context.Context, id int, req dto.UpdateOrderRequest) (dto.OrderResponse, error) {
	return o.repo.UpdateOrder(ctx, id, req)
}

func (o *OrderUsecase) DeleteOrder(ctx context.Context, id int) error {
	return o.repo.DeleteOrder(ctx, id)
}

func (o *OrderUsecase) ListAvailableTables(ctx context.Context) ([]dto.TableResponse, error) {
	return o.repo.ListAvailableTables(ctx)
}

func (o *OrderUsecase) ListPaymentMethods(ctx context.Context) ([]dto.PaymentMethodResponse, error) {
	return o.repo.ListPaymentMethods(ctx)
}
