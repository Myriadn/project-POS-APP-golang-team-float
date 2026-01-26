package repository

import (
	"time"

	"project-POS-APP-golang-team-float/internal/data/entity"
)

// CreateOTP creates a new OTP code in the database
func (r *Repository) CreateOTP(otp *entity.OTPCode) error {
	return r.db.Create(otp).Error
}

// FindValidOTP finds a valid (unused and not expired) OTP for a user
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

// MarkOTPAsUsed marks an OTP as used
func (r *Repository) MarkOTPAsUsed(otpID uint) error {
	return r.db.Model(&entity.OTPCode{}).Where("id = ?", otpID).Update("is_used", true).Error
}

// InvalidatePreviousOTPs marks all previous unused OTPs for a user as used
func (r *Repository) InvalidatePreviousOTPs(userID uint, otpType string) error {
	return r.db.Model(&entity.OTPCode{}).
		Where("user_id = ? AND type = ? AND is_used = ?", userID, otpType, false).
		Update("is_used", true).Error
}
