package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"project-POS-APP-golang-team-float/internal/domain/entity"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create(session *entity.Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepository) FindByToken(token uuid.UUID) (*entity.Session, error) {
	var session entity.Session
	err := r.db.Where("token = ? AND is_active = ?", token, true).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepository) UpdateLastActivity(sessionID uint) error {
	return r.db.Model(&entity.Session{}).Where("id = ?", sessionID).Update("last_activity_at", time.Now()).Error
}

func (r *SessionRepository) Invalidate(token uuid.UUID) error {
	return r.db.Model(&entity.Session{}).Where("token = ?", token).Update("is_active", false).Error
}

func (r *SessionRepository) InvalidateAllByUserID(userID uint) error {
	return r.db.Model(&entity.Session{}).Where("user_id = ?", userID).Update("is_active", false).Error
}

func (r *SessionRepository) DeleteExpired() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&entity.Session{}).Error
}
