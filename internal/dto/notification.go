package dto

import "time"

type NotificationResponse struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Type      string     `json:"type"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	ReadAt    *time.Time `json:"read_at"`
}

type UpdateNotificationStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=new read"`
}

type CreateNotificationRequest struct {
	UserID  uint   `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
	Type    string `json:"type"`
}
