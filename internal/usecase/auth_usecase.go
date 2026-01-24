package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/config"
	"project-POS-APP-golang-team-float/internal/domain/entity"
	"project-POS-APP-golang-team-float/internal/infrastructure/email"
	"project-POS-APP-golang-team-float/internal/repository"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type AuthUsecase struct {
	userRepo    *repository.UserRepository
	sessionRepo *repository.SessionRepository
	otpRepo     *repository.OTPRepository
	emailSvc    *email.SMTPService
	cfg         *config.Config
}

func NewAuthUsecase(
	userRepo *repository.UserRepository,
	sessionRepo *repository.SessionRepository,
	otpRepo *repository.OTPRepository,
	emailSvc *email.SMTPService,
	cfg *config.Config,
) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		otpRepo:     otpRepo,
		emailSvc:    emailSvc,
		cfg:         cfg,
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=4"`
}

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	OTP         string `json:"otp" binding:"required,len=4"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (u *AuthUsecase) Login(req LoginRequest) (*LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is inactive")
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	u.otpRepo.InvalidatePreviousOTPs(user.ID, "login")

	otpCode, err := utils.GenerateOTP(4)
	if err != nil {
		return nil, errors.New("failed to generate OTP")
	}

	otp := &entity.OTPCode{
		UserID:    user.ID,
		Code:      otpCode,
		Type:      "login",
		ExpiresAt: time.Now().Add(time.Duration(u.cfg.OTP.ExpireMinutes) * time.Minute),
	}

	if err := u.otpRepo.Create(otp); err != nil {
		return nil, errors.New("failed to create OTP")
	}

	if err := u.emailSvc.SendOTP(user.Email, otpCode); err != nil {
		return nil, errors.New("failed to send OTP email")
	}

	return &LoginResponse{Message: "OTP sent to your email"}, nil
}

func (u *AuthUsecase) VerifyOTP(req VerifyOTPRequest, ipAddress, userAgent string) (*TokenResponse, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	otp, err := u.otpRepo.FindValidOTP(user.ID, req.OTP, "login")
	if err != nil {
		return nil, errors.New("invalid or expired OTP")
	}

	if err := u.otpRepo.MarkAsUsed(otp.ID); err != nil {
		return nil, errors.New("failed to verify OTP")
	}

	expiresAt := time.Now().Add(time.Duration(u.cfg.Session.ExpireHours) * time.Hour)
	session := &entity.Session{
		UserID:         user.ID,
		Token:          uuid.New(),
		IPAddress:      ipAddress,
		UserAgent:      userAgent,
		IsActive:       true,
		ExpiresAt:      expiresAt,
		LastActivityAt: time.Now(),
	}

	if err := u.sessionRepo.Create(session); err != nil {
		return nil, errors.New("failed to create session")
	}

	return &TokenResponse{
		Token:     session.Token.String(),
		ExpiresAt: expiresAt,
	}, nil
}

func (u *AuthUsecase) CheckEmail(req CheckEmailRequest) (*LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email not found")
	}

	u.otpRepo.InvalidatePreviousOTPs(user.ID, "reset_password")

	otpCode, err := utils.GenerateOTP(4)
	if err != nil {
		return nil, errors.New("failed to generate OTP")
	}

	otp := &entity.OTPCode{
		UserID:    user.ID,
		Code:      otpCode,
		Type:      "reset_password",
		ExpiresAt: time.Now().Add(time.Duration(u.cfg.OTP.ExpireMinutes) * time.Minute),
	}

	if err := u.otpRepo.Create(otp); err != nil {
		return nil, errors.New("failed to create OTP")
	}

	if err := u.emailSvc.SendPasswordResetOTP(user.Email, otpCode); err != nil {
		return nil, errors.New("failed to send OTP email")
	}

	return &LoginResponse{Message: "OTP sent to your email"}, nil
}

func (u *AuthUsecase) ResetPassword(req ResetPasswordRequest) error {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return errors.New("email not found")
	}

	otp, err := u.otpRepo.FindValidOTP(user.ID, req.OTP, "reset_password")
	if err != nil {
		return errors.New("invalid or expired OTP")
	}

	if err := u.otpRepo.MarkAsUsed(otp.ID); err != nil {
		return errors.New("failed to verify OTP")
	}

	passwordHash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	if err := u.userRepo.UpdatePassword(user.ID, passwordHash); err != nil {
		return errors.New("failed to update password")
	}

	u.sessionRepo.InvalidateAllByUserID(user.ID)

	return nil
}

func (u *AuthUsecase) Logout(token uuid.UUID) error {
	return u.sessionRepo.Invalidate(token)
}

func (u *AuthUsecase) ValidateSession(token uuid.UUID) (*entity.User, error) {
	session, err := u.sessionRepo.FindByToken(token)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	if time.Now().After(session.ExpiresAt) {
		u.sessionRepo.Invalidate(token)
		return nil, errors.New("session expired")
	}

	u.sessionRepo.UpdateLastActivity(session.ID)

	user, err := u.userRepo.FindByID(session.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !user.IsActive {
		return nil, errors.New("account is inactive")
	}

	return user, nil
}
