package usecase

import (
	"project-POS-APP-golang-team-float/internal/data/repository"
)

type Usecase struct {
	repo             *repository.Repository
	emailSvc         EmailService
	otpExpireMinutes int
	sessionExpireHrs int
}

type EmailService interface {
	SendOTP(to, otp string) error
	SendPasswordResetOTP(to, otp string) error
}

func NewUsecase(repo *repository.Repository, emailSvc EmailService, otpExpireMinutes, sessionExpireHrs int) *Usecase {
	return &Usecase{
		repo:             repo,
		emailSvc:         emailSvc,
		otpExpireMinutes: otpExpireMinutes,
		sessionExpireHrs: sessionExpireHrs,
	}
}
