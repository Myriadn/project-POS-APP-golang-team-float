package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	SMTP     SMTPConfig
	OTP      OTPConfig
	Session  SessionConfig
}

type AppConfig struct {
	Env  string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type OTPConfig struct {
	ExpireMinutes int
}

type SessionConfig struct {
	ExpireHours int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	smtpPort, _ := strconv.Atoi(getEnv("SMTP_PORT", "587"))
	otpExpire, _ := strconv.Atoi(getEnv("OTP_EXPIRE_MINUTES", "5"))
	sessionExpire, _ := strconv.Atoi(getEnv("SESSION_EXPIRE_HOURS", "24"))

	return &Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "db_posapp"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		SMTP: SMTPConfig{
			Host:     getEnv("SMTP_HOST", "smtp.gmail.com"),
			Port:     smtpPort,
			User:     getEnv("SMTP_USER", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			From:     getEnv("SMTP_FROM", "noreply@posapp.com"),
		},
		OTP: OTPConfig{
			ExpireMinutes: otpExpire,
		},
		Session: SessionConfig{
			ExpireHours: sessionExpire,
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
