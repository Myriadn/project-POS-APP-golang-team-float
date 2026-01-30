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

	// Create shared usecase and middleware
	uc := usecase.NewUsecase(cfg.Repo, cfg.RepoSM, cfg.EmailSvc, cfg.OTPExpireMinutes, cfg.SessionExpireHrs)
	authMw := middleware.NewAuthMiddleware(uc, uc) // uc 2 times because both implemented on middleware

	wireAuth(api, uc, authMw)
	wireDashboard(api, uc, authMw)
	wireStaffManagement(api, uc, authMw)
	wireOrders(api, uc, authMw)

	return router
}

func wireOrders(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	orderAdaptor := adaptor.NewOrderAdaptor(uc.OrderUsecase)

	orders := router.Group("/orders")
	orders.Use(authMw.Authenticate())
	{
		orders.GET("", orderAdaptor.ListOrders)
		orders.POST("", orderAdaptor.CreateOrder)
		orders.PUT(":id", orderAdaptor.UpdateOrder)
		orders.DELETE(":id", orderAdaptor.DeleteOrder)
	}

	router.GET("/tables/available", authMw.Authenticate(), orderAdaptor.ListAvailableTables)
	router.GET("/payment-methods", authMw.Authenticate(), orderAdaptor.ListPaymentMethods)
}

func wireAuth(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	authAdaptor := adaptor.NewAuthAdaptor(uc)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authAdaptor.Login)
		auth.POST("/verify-otp", authAdaptor.VerifyOTP)
		auth.POST("/check-email", authAdaptor.CheckEmail)
		auth.POST("/reset-password", authAdaptor.ResetPassword)
		auth.POST("/logout", authMw.Authenticate(), authAdaptor.Logout)

	}

}

func wireDashboard(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	dashboardAdaptor := adaptor.NewDashboardAdaptor(uc)

	dashboard := router.Group("/dashboard")
	dashboard.Use(authMw.Authenticate())
	// dashboard.Use(authMw.RequirePermission("view_dashboard"))
	{
		dashboard.GET("", dashboardAdaptor.GetDashboardSummary)
		dashboard.GET("/daily-sales", dashboardAdaptor.GetDailySales)
		dashboard.GET("/monthly-sales", dashboardAdaptor.GetMonthlySales)
		dashboard.GET("/tables", dashboardAdaptor.GetTableSummary)
		dashboard.GET("/popular-products", dashboardAdaptor.GetPopularProducts)
		dashboard.GET("/new-products", dashboardAdaptor.GetNewProducts)
	}
}

func wireStaffManagement(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	staffManagementAdaptor := adaptor.NewStaffManagementAdaptor(uc.StaffManagementUsecase)

	staffManagement := router.Group("/staff-management")
	staffManagement.Use(authMw.Authenticate())
	staffManagement.Use(authMw.RequirePermission("manage_staff"))
	{
		staffManagement.POST("/create", staffManagementAdaptor.CreateNewStaffManagement)
		staffManagement.PATCH("/update/:id", staffManagementAdaptor.UpdateStaffManagement)
		staffManagement.GET("/:id", staffManagementAdaptor.GetDetailStaffManagement)
		staffManagement.GET("", staffManagementAdaptor.GetAllStaffManagement)
		staffManagement.DELETE("/delete/:id", staffManagementAdaptor.DeleteStaffManagement)
	}
}
