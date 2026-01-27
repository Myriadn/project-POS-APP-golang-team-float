package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"project-POS-APP-golang-team-float/internal/data/entity"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) DB() *gorm.DB {
	return r.db
}

// User Repository Methods
func (r *Repository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindUserByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *Repository) UpdatePassword(userID uint, passwordHash string) error {
	return r.db.Model(&entity.User{}).Where("id = ?", userID).Update("password_hash", passwordHash).Error
}

func (r *Repository) EmailExists(email string) bool {
	var count int64
	r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// Session Repository Methods
func (r *Repository) CreateSession(session *entity.Session) error {
	return r.db.Create(session).Error
}

func (r *Repository) FindSessionByToken(token uuid.UUID) (*entity.Session, error) {
	var session entity.Session
	err := r.db.Where("token = ? AND is_active = ?", token, true).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *Repository) UpdateSessionActivity(sessionID uint) error {
	return r.db.Model(&entity.Session{}).Where("id = ?", sessionID).Update("last_activity_at", time.Now()).Error
}

func (r *Repository) InvalidateSession(token uuid.UUID) error {
	return r.db.Model(&entity.Session{}).Where("token = ?", token).Update("is_active", false).Error
}

func (r *Repository) InvalidateAllUserSessions(userID uint) error {
	return r.db.Model(&entity.Session{}).Where("user_id = ?", userID).Update("is_active", false).Error
}

// OTP Repository Methods
func (r *Repository) CreateOTP(otp *entity.OTPCode) error {
	return r.db.Create(otp).Error
}

func (r *Repository) FindValidOTP(userID uint, code, otpType string) (*entity.OTPCode, error) {
	var otp entity.OTPCode
	err := r.db.Where(
		"user_id = ? AND code = ? AND type = ? AND is_used = ? AND expires_at > ?",
		userID, code, otpType, false, time.Now(),
	).First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *Repository) MarkOTPAsUsed(otpID uint) error {
	return r.db.Model(&entity.OTPCode{}).Where("id = ?", otpID).Update("is_used", true).Error
}

func (r *Repository) InvalidatePreviousOTPs(userID uint, otpType string) error {
	return r.db.Model(&entity.OTPCode{}).
		Where("user_id = ? AND type = ? AND is_used = ?", userID, otpType, false).
		Update("is_used", true).Error
}

// pengecekan permission berdarsarkan id dan code nya
func (r *Repository) Allowed(userID uint, code string) (bool, error) {
	var exists bool

	query := `
        SELECT EXISTS (
            SELECT 1
            FROM users u
            JOIN role_permissions rp ON rp.role_id = u.role_id
            JOIN permissions p ON p.id = rp.permission_id
            WHERE u.id = ? AND p.code = ?
        )
    `
	err := r.db.Raw(query, userID, code).Scan(&exists).Error

	if err != nil {
		return false, err
	}

	return exists, nil
}

// menemukan id user dari session Token
func (r *Repository) FindUserIDBySession(sessionToken string) (uint, error) {
	var userID uint
	err := r.db.Select("user_id").Where("token=?", sessionToken).First(&userID).Error
	if err != nil {
		return 0, err
	}
	return userID, nil
}
