package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type SMTPService struct {
	dialer *gomail.Dialer
	from   string
}

func NewSMTPService(cfg SMTPConfig) *SMTPService {
	dialer := gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Password)
	return &SMTPService{
		dialer: dialer,
		from:   cfg.From,
	}
}

func (s *SMTPService) SendOTP(to, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "COSYPOS - Your OTP Code")
	m.SetBody("text/html", fmt.Sprintf(`
		<h2>Your OTP Code</h2>
		<p>Use this code to complete your authentication:</p>
		<h1 style="color: #E91E63; letter-spacing: 5px;">%s</h1>
		<p>This code will expire in 5 minutes.</p>
		<p>If you did not request this code, please ignore this email.</p>
	`, otp))

	return s.dialer.DialAndSend(m)
}

func (s *SMTPService) SendPasswordResetOTP(to, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "COSYPOS - Password Reset OTP")
	m.SetBody("text/html", fmt.Sprintf(`
		<h2>Password Reset Request</h2>
		<p>Use this code to reset your password:</p>
		<h1 style="color: #E91E63; letter-spacing: 5px;">%s</h1>
		<p>This code will expire in 5 minutes.</p>
		<p>If you did not request this, please ignore this email.</p>
	`, otp))

	return s.dialer.DialAndSend(m)
}
