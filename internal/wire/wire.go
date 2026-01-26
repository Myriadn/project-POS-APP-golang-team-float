package wire

import (
	"github.com/gin-gonic/gin"

	"project-POS-APP-golang-team-float/internal/adaptor"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/email"
	"project-POS-APP-golang-team-float/pkg/middleware"
)

type WireConfig struct {
	Repo             *repository.Repository
	RepoSM           repository.StaffManagementRepoInterface
	EmailSvc         *email.SMTPService
	OTPExpireMinutes int
	SessionExpireHrs int
}

func Wiring(cfg WireConfig) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	wireAuth(api, cfg)

	return router
}

func wireAuth(router *gin.RouterGroup, cfg WireConfig) {
	uc := usecase.NewUsecase(cfg.Repo, cfg.RepoSM, cfg.EmailSvc, cfg.OTPExpireMinutes, cfg.SessionExpireHrs)
	authMw := middleware.NewAuthMiddleware(uc)
	authAdaptor := adaptor.NewAuthAdaptor(uc)
	staffManagementAdaptor := adaptor.NewStaffManagementAdaptor(uc.StaffManagementUsecase)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authAdaptor.Login)
		auth.POST("/verify-otp", authAdaptor.VerifyOTP)
		auth.POST("/check-email", authAdaptor.CheckEmail)
		auth.POST("/reset-password", authAdaptor.ResetPassword)
		auth.POST("/logout", authMw.Authenticate(), authAdaptor.Logout)

	}
	staffManagement := router.Group("/staff-management")
	{
		staffManagement.POST("/create", staffManagementAdaptor.CreateNewStaffManagement)
	}
}
