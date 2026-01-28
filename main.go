package main

import (
	"project-POS-APP-golang-team-float/cmd"
	"project-POS-APP-golang-team-float/internal/data"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/wire"
	"project-POS-APP-golang-team-float/pkg/database"
	"project-POS-APP-golang-team-float/pkg/email"
	"project-POS-APP-golang-team-float/pkg/logger"
	"project-POS-APP-golang-team-float/pkg/utils"

	"go.uber.org/zap"
)

func main() {
	// Load Config
	cfg, err := utils.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	// Init Logger
	if err := logger.Init(cfg.App.Env); err != nil {
		panic("Failed to init logger: " + err.Error())
	}
	defer logger.Sync()

	logger.Info("Starting POS App Server", zap.String("env", cfg.App.Env))

	// Init Database
	db, err := database.Connect(database.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Name:     cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode,
	})
	if err != nil {
		logger.Fatal("Failed to connect database", zap.Error(err))
	}
	logger.Info("Database connected")

	// AutoMigrate
	if err := data.Migrate(db); err != nil {
		logger.Fatal("Failed to migrate database", zap.Error(err))
	}
	logger.Info("Database migrated and triggers set")

	// Seed Data
	if err := data.Seed(db); err != nil {
		logger.Fatal("Failed to seed database", zap.Error(err))
	}
	logger.Info("Database seeded")

	// Dependency Injection - Repository
	repo := repository.NewRepository(db)
	repoSM := repository.NewStaffManagementRepo(db)
	repoCategory := repository.NewCategoryMenuRepo(db)

	// Email Service
	emailSvc := email.NewSMTPService(email.SMTPConfig{
		Host:     cfg.SMTP.Host,
		Port:     cfg.SMTP.Port,
		User:     cfg.SMTP.User,
		Password: cfg.SMTP.Password,
		From:     cfg.SMTP.From,
	})

	// Wiring
	router := wire.Wiring(wire.WireConfig{
		Repo:             repo,
		RepoSM:           repoSM,
		Category:         repoCategory,
		EmailSvc:         emailSvc,
		OTPExpireMinutes: cfg.OTP.ExpireMinutes,
		SessionExpireHrs: cfg.Session.ExpireHours,
	})

	// Start Server
	cmd.APIServer(router, cfg.App.Port)
}
