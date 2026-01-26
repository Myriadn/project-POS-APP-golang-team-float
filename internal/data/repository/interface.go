package repository

import (
	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/internal/data/entity"
)

// RepositoryInterface defines all repository methods for dependency injection and testing
type RepositoryInterface interface {
	UserRepository
	SessionRepository
	OTPRepository
	RoleRepository
}

// UserRepository defines user-related database operations
type UserRepository interface {
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByID(id uint) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	UpdatePassword(userID uint, passwordHash string) error
	EmailExists(email string) bool
}

// SessionRepository defines session-related database operations
type SessionRepository interface {
	CreateSession(session *entity.Session) error
	FindSessionByToken(token uuid.UUID) (*entity.Session, error)
	UpdateSessionActivity(sessionID uint) error
	InvalidateSession(token uuid.UUID) error
	InvalidateAllUserSessions(userID uint) error
}

// OTPRepository defines OTP-related database operations
type OTPRepository interface {
	CreateOTP(otp *entity.OTPCode) error
	FindValidOTP(userID uint, code, otpType string) (*entity.OTPCode, error)
	MarkOTPAsUsed(otpID uint) error
	InvalidatePreviousOTPs(userID uint, otpType string) error
}

// RoleRepository defines role-related database operations
type RoleRepository interface {
	FindRoleByName(name string) (*entity.Role, error)
	FindRoleByID(id uint) (*entity.Role, error)
	GetAllRoles() ([]entity.Role, error)
	CreateRole(role *entity.Role) error
	UpdateRole(role *entity.Role) error
	DeleteRole(id uint) error
}

// Ensure Repository implements RepositoryInterface
var _ RepositoryInterface = (*Repository)(nil)
