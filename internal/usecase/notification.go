package usecase

import (
	"context"

	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
)

type NotificationUsecaseInterface interface {
	ListNotifications(ctx context.Context, userID uint) ([]dto.NotificationResponse, error)
	CreateNotification(ctx context.Context, req dto.CreateNotificationRequest) error
	UpdateNotificationStatus(ctx context.Context, id uint, status string) error
	DeleteNotification(ctx context.Context, id uint) error
}

type NotificationUsecase struct {
	repo repository.NotificationRepository
}

func NewNotificationUsecase(repo repository.NotificationRepository) NotificationUsecaseInterface {
	return &NotificationUsecase{repo: repo}
}

func (u *NotificationUsecase) ListNotifications(ctx context.Context, userID uint) ([]dto.NotificationResponse, error) {
	notifs, err := u.repo.ListNotifications(ctx, userID)
	if err != nil {
		return nil, err
	}
	var result []dto.NotificationResponse
	for _, n := range notifs {
		result = append(result, dto.NotificationResponse{
			ID:        n.ID,
			UserID:    n.UserID,
			Title:     n.Title,
			Message:   n.Message,
			Type:      n.Type,
			Status:    n.Status,
			CreatedAt: n.CreatedAt,
			ReadAt:    n.ReadAt,
		})
	}
	return result, nil
}

func (u *NotificationUsecase) CreateNotification(ctx context.Context, req dto.CreateNotificationRequest) error {
	notif := &entity.Notification{
		UserID:  req.UserID,
		Title:   req.Title,
		Message: req.Message,
		Type:    req.Type,
		Status:  "new",
	}
	return u.repo.CreateNotification(ctx, notif)
}

func (u *NotificationUsecase) UpdateNotificationStatus(ctx context.Context, id uint, status string) error {
	return u.repo.UpdateNotificationStatus(ctx, id, status)
}

func (u *NotificationUsecase) DeleteNotification(ctx context.Context, id uint) error {
	return u.repo.DeleteNotification(ctx, id)
}
