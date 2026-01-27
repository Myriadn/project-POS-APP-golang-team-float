package repository

import (
	"time"

	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/internal/data/entity"
)

// CreateSession creates a new session in the database
func (r *Repository) CreateSession(session *entity.Session) error {
	return r.db.Create(session).Error
}

// FindSessionByToken finds an active session by its token
func (r *Repository) FindSessionByToken(token uuid.UUID) (*entity.Session, error) {
	var session entity.Session
	err := r.db.Where("token = ? AND is_active = ?", token, true).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// UpdateSessionActivity updates the last activity timestamp for a session
func (r *Repository) UpdateSessionActivity(sessionID uint) error {
	return r.db.Model(&entity.Session{}).Where("id = ?", sessionID).Update("last_activity_at", time.Now()).Error
}

// InvalidateSession invalidates a session by its token
func (r *Repository) InvalidateSession(token uuid.UUID) error {
	return r.db.Model(&entity.Session{}).Where("token = ?", token).Update("is_active", false).Error
}

// InvalidateAllUserSessions invalidates all sessions for a specific user
func (r *Repository) InvalidateAllUserSessions(userID uint) error {
	return r.db.Model(&entity.Session{}).Where("user_id = ?", userID).Update("is_active", false).Error
}
