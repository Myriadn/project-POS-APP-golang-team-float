package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	ListNotifications(ctx context.Context, userID uint) ([]entity.Notification, error)
	CreateNotification(ctx context.Context, notif *entity.Notification) error
	UpdateNotificationStatus(ctx context.Context, id uint, status string) error
	DeleteNotification(ctx context.Context, id uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) ListNotifications(ctx context.Context, userID uint) ([]entity.Notification, error) {
	var notifications []entity.Notification
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepository) CreateNotification(ctx context.Context, notif *entity.Notification) error {
	return r.db.WithContext(ctx).Create(notif).Error
}

func (r *notificationRepository) UpdateNotificationStatus(ctx context.Context, id uint, status string) error {
	update := map[string]interface{}{"status": status}
	if status == "read" {
		update["read_at"] = gorm.Expr("NOW()")
	} else {
		update["read_at"] = nil
	}
	return r.db.WithContext(ctx).Model(&entity.Notification{}).Where("id = ?", id).Updates(update).Error
}

func (r *notificationRepository) DeleteNotification(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Notification{}, id).Error
}
