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
	OrderRepo        repository.OrderRepository
	ReservationRepo  repository.ReservationRepository
	NotificationRepo repository.NotificationRepository
	Category         repository.CategoryMenuRepoInterface
	Product          repository.ProductMenuRepoInterface
	Profile          repository.ProfileRepoInterface
	ReportRepo       repository.ReportRepoInterface
	EmailSvc         *email.SMTPService
	OTPExpireMinutes int
	SessionExpireHrs int
}

func Wiring(cfg WireConfig) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	// Create shared usecase and middleware
	uc := usecase.NewUsecase(cfg.Repo, cfg.RepoSM, cfg.OrderRepo, cfg.ReservationRepo, cfg.NotificationRepo, cfg.Category, cfg.Product, cfg.Profile, cfg.ReportRepo, cfg.EmailSvc, cfg.OTPExpireMinutes, cfg.SessionExpireHrs)
	authMw := middleware.NewAuthMiddleware(uc, uc)

	wireAuth(api, uc, authMw)
	wireDashboard(api, uc, authMw)
	wireStaffManagement(api, uc, authMw)
	wireCategoryMenu(api, uc, authMw)
	wireProductMenu(api, uc, authMw)
	wireReport(api, uc, authMw)
	wireProfile(api, uc, authMw)
	wireOrders(api, uc, authMw)
	wireReservations(api, uc, authMw)
	wireNotifications(api, uc, authMw)

	return router
}

func wireNotifications(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
       notificationAdaptor := adaptor.NewNotificationAdaptor(uc.NotificationUsecase)

       notifications := router.Group("/notifications")
       notifications.Use(authMw.Authenticate())
       {
	       notifications.GET("", notificationAdaptor.ListNotifications)
	       notifications.PATCH(":id", notificationAdaptor.UpdateNotificationStatus)
	       notifications.DELETE(":id", notificationAdaptor.DeleteNotification)
       }
}


func wireReservations(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	reservationAdaptor := adaptor.NewReservationAdaptor(uc.ReservationUsecase)

	reservations := router.Group("/reservations")
	reservations.Use(authMw.Authenticate())
	{
		reservations.GET("", reservationAdaptor.ListReservations)
		reservations.GET(":id", reservationAdaptor.GetReservationByID)
		reservations.POST("", reservationAdaptor.CreateReservation)
		reservations.PATCH(":id", reservationAdaptor.UpdateReservation)
	}
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
	{
		staffManagement.POST("/create", authMw.RequirePermission("user:create"), staffManagementAdaptor.CreateNewStaffManagement)
		staffManagement.PATCH("/update/:id", authMw.RequirePermission("user:update"), staffManagementAdaptor.UpdateStaffManagement)
		staffManagement.GET("/:id", authMw.RequirePermission("user:read"), staffManagementAdaptor.GetDetailStaffManagement)
		staffManagement.GET("", authMw.RequirePermission("user:read"), staffManagementAdaptor.GetAllStaffManagement)
		staffManagement.DELETE("/delete/:id", authMw.RequirePermission("user:delete"), staffManagementAdaptor.DeleteStaffManagement)
	}
}

func wireCategoryMenu(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	CategoryMenuAdaptor := adaptor.NewCategoryMenuAdaptor(uc.CategoryMenuUsecase)

	CategoryMenu := router.Group("/category-menu")
	CategoryMenu.Use(authMw.Authenticate())
	{
		CategoryMenu.POST("/create", authMw.RequirePermission("category:create"), CategoryMenuAdaptor.CreateNewCategoryMenu)
		CategoryMenu.PATCH("/update/:id", authMw.RequirePermission("category:update"), CategoryMenuAdaptor.UpdateCategoryMenu)
		CategoryMenu.GET("/:id", authMw.RequirePermission("category:read"), CategoryMenuAdaptor.GetDetailCategoryMenu)
		CategoryMenu.GET("", authMw.RequirePermission("category:read"), CategoryMenuAdaptor.GetAllCategoryMenu)
		CategoryMenu.DELETE("/delete/:id", authMw.RequirePermission("category:delete"), CategoryMenuAdaptor.DeleteCategoryMenu)

	}
}

func wireProductMenu(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	ProductMenuAdaptor := adaptor.NewProductMenuAdaptor(uc.ProductMenuUsecase)

	ProductMenu := router.Group("/product-menu")
	ProductMenu.Use(authMw.Authenticate())
	{
		ProductMenu.POST("/create", authMw.RequirePermission("product:create"), ProductMenuAdaptor.CreateNewProductMenu)
		ProductMenu.PATCH("/update/:id", authMw.RequirePermission("product:update"), ProductMenuAdaptor.UpdateProductMenu)
		ProductMenu.GET("/:id", authMw.RequirePermission("product:read"), ProductMenuAdaptor.GetDetailProductMenu)
		ProductMenu.GET("", authMw.RequirePermission("product:read"), ProductMenuAdaptor.GetAllStaffProductMenu)
		ProductMenu.DELETE("/delete/:id", authMw.RequirePermission("product:delete"), ProductMenuAdaptor.DeleteProductMenu)

	}
}

func wireReport(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	reportAdaptor := adaptor.NewReportAdaptor(uc.ReportUsecase)

	reports := router.Group("/reports")
	reports.Use(authMw.Authenticate())
	{

		reports.GET("/revenue", reportAdaptor.GetRevenueReport)
	}
}

func wireProfile(router *gin.RouterGroup, uc *usecase.Usecase, authMw *middleware.AuthMiddleware) {
	ProfileAdaptor := adaptor.NewProfileAdaptor(uc.ProfileUsecase)

	Profile := router.Group("/profile")
	Profile.Use(authMw.Authenticate())
	{
		Profile.PATCH("/update", ProfileAdaptor.UpdateProfile)
	}
	ManageAccsess := router.Group("/manage-accsess")
	ManageAccsess.Use(authMw.Authenticate())
	{
		ManageAccsess.GET("/admin", ProfileAdaptor.GetAllAdminUser)
	}
}
