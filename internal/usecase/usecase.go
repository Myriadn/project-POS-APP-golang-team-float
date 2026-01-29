package usecase

import (
	"project-POS-APP-golang-team-float/internal/data/repository"
)

// Usecase handles business logic for the application
type Usecase struct {
	repo                   *repository.Repository
	StaffManagementUsecase StaffManagementUsecaseInterface
	CategoryMenuUsecase    CategoryMenuUsecaseInterface
	ProductMenuUsecase     ProductMenuUsecaseInterface
	ReportUsecase          ReportUsecaseInterface
	emailSvc               EmailService
	otpExpireMinutes       int
	sessionExpireHrs       int
}

type EmailService interface {
	SendOTP(to, otp string) error
	SendPasswordResetOTP(to, otp string) error
}

func NewUsecase(repo *repository.Repository, repoSM repository.StaffManagementRepoInterface, Category repository.CategoryMenuRepoInterface, product repository.ProductMenuRepoInterface, report repository.ReportRepoInterface, emailSvc EmailService, otpExpireMinutes, sessionExpireHrs int) *Usecase {
	return &Usecase{
		repo:                   repo,
		emailSvc:               emailSvc,
		otpExpireMinutes:       otpExpireMinutes,
		sessionExpireHrs:       sessionExpireHrs,
		StaffManagementUsecase: NewStaffManagementUsecase(repoSM),
		CategoryMenuUsecase:    NewCategoryMenuUsecase(Category),
		ProductMenuUsecase:     NewProductMenuUsecase(product),
		ReportUsecase:          NewReportUsecase(report),
	}
}

func (u *Usecase) GetRepository() repository.RepositoryInterface {
	return u.repo
}
