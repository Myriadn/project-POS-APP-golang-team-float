package repository

import (
	"time"

	"gorm.io/gorm"

	"project-POS-APP-golang-team-float/internal/domain/entity"
)

type OTPRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) *OTPRepository {
	return &OTPRepository{db: db}
}

func (r *OTPRepository) Create(otp *entity.OTPCode) error {
	return r.db.Create(otp).Error
}

func (r *OTPRepository) FindValidOTP(userID uint, code, otpType string) (*entity.OTPCode, error) {
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

func (r *OTPRepository) MarkAsUsed(otpID uint) error {
	return r.db.Model(&entity.OTPCode{}).Where("id = ?", otpID).Update("is_used", true).Error
}

func (r *OTPRepository) InvalidatePreviousOTPs(userID uint, otpType string) error {
	return r.db.Model(&entity.OTPCode{}).
		Where("user_id = ? AND type = ? AND is_used = ?", userID, otpType, false).
		Update("is_used", true).Error
}

func (r *OTPRepository) DeleteExpired() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&entity.OTPCode{}).Error
}
