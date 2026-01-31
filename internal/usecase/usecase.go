package usecase

import (
	"project-POS-APP-golang-team-float/internal/data/repository"
)

// Usecase handles business logic for the application
type Usecase struct {
	repo                   *repository.Repository
	StaffManagementUsecase StaffManagementUsecaseInterface
	OrderUsecase           OrderUsecaseInterface
	CategoryMenuUsecase    CategoryMenuUsecaseInterface
	ProductMenuUsecase     ProductMenuUsecaseInterface
	ProfileUsecase         ProfileUsecaseInterface
	ReportUsecase          ReportUsecaseInterface
	ReservationUsecase     ReservationUsecaseInterface
	NotificationUsecase    NotificationUsecaseInterface
	emailSvc               EmailService
	otpExpireMinutes       int
	sessionExpireHrs       int
}

type EmailService interface {
	SendOTP(to, otp string) error
	SendPasswordResetOTP(to, otp string) error
}

// Dihapus: fungsi NewUsecase yang tidak lengkap/duplikat
func NewUsecase(
	repo *repository.Repository,
	repoSM repository.StaffManagementRepoInterface,
	orderRepo repository.OrderRepository,
	reservationRepo repository.ReservationRepository,
	notificationRepo repository.NotificationRepository,
	Category repository.CategoryMenuRepoInterface,
	product repository.ProductMenuRepoInterface,
	profile repository.ProfileRepoInterface,
	report repository.ReportRepoInterface,
	emailSvc EmailService,
	otpExpireMinutes, sessionExpireHrs int,
) *Usecase {
	notifUC := NewNotificationUsecase(notificationRepo)
	return &Usecase{
		  repo:                   repo,
		  emailSvc:               emailSvc,
		  otpExpireMinutes:       otpExpireMinutes,
		  sessionExpireHrs:       sessionExpireHrs,
		  StaffManagementUsecase: NewStaffManagementUsecase(repoSM),
		  OrderUsecase:           NewOrderUsecase(orderRepo),
		  ReservationUsecase:     NewReservationUsecase(reservationRepo),
		  NotificationUsecase:    notifUC,
		  CategoryMenuUsecase:    NewCategoryMenuUsecase(Category),
		  ProductMenuUsecase:     NewProductMenuUsecase(product),
		  ProfileUsecase:         NewProfileUsecase(profile),
		  ReportUsecase:          NewReportUsecase(report),
	}
}

func (u *Usecase) GetRepository() repository.RepositoryInterface {
	return u.repo
}
