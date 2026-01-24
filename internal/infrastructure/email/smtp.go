package email

import (
	"fmt"

	"gopkg.in/gomail.v2"

	"project-POS-APP-golang-team-float/config"
)

type SMTPService struct {
	dialer *gomail.Dialer
	from   string
}

func NewSMTPService(cfg config.SMTPConfig) *SMTPService {
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

func (s *SMTPService) SendNewAdminPassword(to, name, password string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "COSYPOS - Your Admin Account")
	m.SetBody("text/html", fmt.Sprintf(`
		<h2>Welcome to COSYPOS, %s!</h2>
		<p>Your admin account has been created. Here are your login credentials:</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Password:</strong> %s</p>
		<p>Please change your password after your first login.</p>
	`, name, to, password))

	return s.dialer.DialAndSend(m)
}
