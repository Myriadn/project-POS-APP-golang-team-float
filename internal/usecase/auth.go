package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/pkg/utils"
)

func (u *Usecase) Login(req dto.LoginRequest) (*dto.MessageResponse, error) {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is inactive")
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	u.repo.InvalidatePreviousOTPs(user.ID, "login")

	otpCode, err := utils.GenerateOTP(4)
	if err != nil {
		return nil, errors.New("failed to generate OTP")
	}

	otp := &entity.OTPCode{
		UserID:    user.ID,
		Code:      otpCode,
		Type:      "login",
		ExpiresAt: time.Now().Add(time.Duration(u.otpExpireMinutes) * time.Minute),
	}

	if err := u.repo.CreateOTP(otp); err != nil {
		return nil, errors.New("failed to create OTP")
	}

	if err := u.emailSvc.SendOTP(user.Email, otpCode); err != nil {
		return nil, errors.New("failed to send OTP email")
	}

	return &dto.MessageResponse{Message: "OTP sent to your email"}, nil
}

func (u *Usecase) VerifyOTP(req dto.VerifyOTPRequest, ipAddress, userAgent string) (*dto.TokenResponse, error) {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	otp, err := u.repo.FindValidOTP(user.ID, req.OTP, "login")
	if err != nil {
		return nil, errors.New("invalid or expired OTP")
	}

	if err := u.repo.MarkOTPAsUsed(otp.ID); err != nil {
		return nil, errors.New("failed to verify OTP")
	}

	expiresAt := time.Now().Add(time.Duration(u.sessionExpireHrs) * time.Hour)
	session := &entity.Session{
		UserID:         user.ID,
		Token:          uuid.New(),
		IPAddress:      ipAddress,
		UserAgent:      userAgent,
		IsActive:       true,
		ExpiresAt:      expiresAt,
		LastActivityAt: time.Now(),
	}

	if err := u.repo.CreateSession(session); err != nil {
		return nil, errors.New("failed to create session")
	}

	return &dto.TokenResponse{
		Token:     session.Token.String(),
		ExpiresAt: expiresAt,
	}, nil
}

func (u *Usecase) CheckEmail(req dto.CheckEmailRequest) (*dto.MessageResponse, error) {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email not found")
	}

	u.repo.InvalidatePreviousOTPs(user.ID, "reset_password")

	otpCode, err := utils.GenerateOTP(4)
	if err != nil {
		return nil, errors.New("failed to generate OTP")
	}

	otp := &entity.OTPCode{
		UserID:    user.ID,
		Code:      otpCode,
		Type:      "reset_password",
		ExpiresAt: time.Now().Add(time.Duration(u.otpExpireMinutes) * time.Minute),
	}

	if err := u.repo.CreateOTP(otp); err != nil {
		return nil, errors.New("failed to create OTP")
	}

	if err := u.emailSvc.SendPasswordResetOTP(user.Email, otpCode); err != nil {
		return nil, errors.New("failed to send OTP email")
	}

	return &dto.MessageResponse{Message: "OTP sent to your email"}, nil
}

func (u *Usecase) ResetPassword(req dto.ResetPasswordRequest) error {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return errors.New("email not found")
	}

	otp, err := u.repo.FindValidOTP(user.ID, req.OTP, "reset_password")
	if err != nil {
		return errors.New("invalid or expired OTP")
	}

	if err := u.repo.MarkOTPAsUsed(otp.ID); err != nil {
		return errors.New("failed to verify OTP")
	}

	passwordHash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	if err := u.repo.UpdatePassword(user.ID, passwordHash); err != nil {
		return errors.New("failed to update password")
	}

	u.repo.InvalidateAllUserSessions(user.ID)
	return nil
}

func (u *Usecase) Logout(token uuid.UUID) error {
	return u.repo.InvalidateSession(token)
}

func (u *Usecase) ValidateSession(token uuid.UUID) (*entity.User, error) {
	session, err := u.repo.FindSessionByToken(token)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	if time.Now().After(session.ExpiresAt) {
		u.repo.InvalidateSession(token)
		return nil, errors.New("session expired")
	}

	u.repo.UpdateSessionActivity(session.ID)

	user, err := u.repo.FindUserByID(session.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !user.IsActive {
		return nil, errors.New("account is inactive")
	}

	return user, nil
}

// pengecekan izin
func (u *Usecase) Allowed(userID uint, code string) (bool, error) {
	allowed, err := u.repo.Allowed(userID, code)
	if err != nil {
		return false, err
	}

	return allowed, nil
}
