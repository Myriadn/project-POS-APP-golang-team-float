package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type SessionValidator interface {
	ValidateSession(token uuid.UUID) (*entity.User, error)
}

type PermissionChecker interface {
	Allowed(userID uint, code string) (bool, error)
}

type AuthMiddleware struct {
	validator   SessionValidator
	permChecker PermissionChecker
}

func NewAuthMiddleware(validator SessionValidator, permChecker PermissionChecker) *AuthMiddleware {
	return &AuthMiddleware{validator: validator, permChecker: permChecker}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string

		cookieToken, err := c.Cookie("session_token")
		if err == nil && cookieToken != "" {
			tokenStr = cookieToken
		} else {
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" && len(authHeader) > 7 {
				tokenStr = authHeader[7:] // Hapus "Bearer "
			}
		}

		if tokenStr == "" {
			utils.Unauthorized(c, "Session token required")
			c.Abort()
			return
		}

		token, err := uuid.Parse(tokenStr)
		if err != nil {
			utils.Unauthorized(c, "Invalid token format")
			c.Abort()
			return
		}

		user, err := m.validator.ValidateSession(token)
		if err != nil {
			c.SetCookie("session_token", "", -1, "/", "localhost", false, true)
			utils.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Set("role", user.Role.Name)
		c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			utils.Unauthorized(c, "Role not found")
			c.Abort()
			return
		}

		roleStr := userRole.(string)

		if slices.Contains(roles, roleStr) {
			c.Next()
			return
		}

		utils.Unauthorized(c, "Insufficient permissions")
		c.Abort()
	}
}

// untuk melihat apakah di izinkan atau tidak
func (m *AuthMiddleware) RequirePermission(permissionCode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil User ID dari Context (ini diset oleh Authenticate sebelumnya)
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			utils.Unauthorized(c, "User ID not found in context")
			c.Abort()
			return
		}

		// Pastikan tipe datanya uint
		userID, ok := userIDInterface.(uint)
		if !ok {
			utils.Unauthorized(c, "Invalid user ID format")
			c.Abort()
			return
		}

		// Panggil Method Allowed di Usecase
		allowed, err := m.permChecker.Allowed(userID, permissionCode)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Error checking permission", err.Error())
			c.Abort()
			return
		}

		if !allowed {
			utils.ErrorResponse(c, http.StatusForbidden, "You don't have permission: "+permissionCode, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
