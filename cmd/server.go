package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"project-POS-APP-golang-team-float/config"
	"project-POS-APP-golang-team-float/internal/delivery/http/handler"
	"project-POS-APP-golang-team-float/internal/delivery/http/middleware"
	"project-POS-APP-golang-team-float/internal/delivery/http/router"
	"project-POS-APP-golang-team-float/internal/domain/entity"
	"project-POS-APP-golang-team-float/internal/infrastructure/database"
	"project-POS-APP-golang-team-float/internal/infrastructure/email"
	"project-POS-APP-golang-team-float/internal/infrastructure/logger"
	"project-POS-APP-golang-team-float/internal/repository"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/migrations"
)

func StartServer() {
	cfg, err := config.Load()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	if err := logger.Init(cfg.App.Env); err != nil {
		panic("Failed to init logger: " + err.Error())
	}
	defer logger.Sync()

	logger.Info("Starting COSYPOS API Server", zap.String("env", cfg.App.Env))

	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect database", zap.Error(err))
	}
	logger.Info("Database connected")

	if err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.OTPCode{},
		&entity.Session{},
		&entity.Category{},
		&entity.Table{},
		&entity.PaymentMethod{},
	); err != nil {
		logger.Fatal("Failed to migrate database", zap.Error(err))
	}
	logger.Info("Database migrated")

	if err := migrations.Seed(db); err != nil {
		logger.Fatal("Failed to seed database", zap.Error(err))
	}
	logger.Info("Database seeded")

	emailSvc := email.NewSMTPService(cfg.SMTP)

	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	otpRepo := repository.NewOTPRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, sessionRepo, otpRepo, emailSvc, cfg)

	authHandler := handler.NewAuthHandler(authUsecase)
	authMiddleware := middleware.NewAuthMiddleware(authUsecase)

	r := router.NewRouter(authHandler, authMiddleware)
	engine := r.Setup()

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: engine,
	}

	go func() {
		logger.Info("Server starting", zap.String("port", cfg.App.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exited")
}
