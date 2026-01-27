package usecase

import (
	"project-POS-APP-golang-team-float/internal/data/repository"
)

// Usecase handles business logic for the application
type Usecase struct {
	repo                   *repository.Repository
	StaffManagementUsecase StaffManagementUsecaseInterface
	emailSvc               EmailService
	otpExpireMinutes       int
	sessionExpireHrs       int
}

type EmailService interface {
	SendOTP(to, otp string) error
	SendPasswordResetOTP(to, otp string) error
}

func NewUsecase(repo *repository.Repository, repoSM repository.StaffManagementRepoInterface, emailSvc EmailService, otpExpireMinutes, sessionExpireHrs int) *Usecase {
	return &Usecase{
		repo:                   repo,
		emailSvc:               emailSvc,
		otpExpireMinutes:       otpExpireMinutes,
		sessionExpireHrs:       sessionExpireHrs,
		StaffManagementUsecase: NewStaffManagementUsecase(repoSM),
	}
}

func (u *Usecase) GetRepository() repository.RepositoryInterface {
	return u.repo
}
