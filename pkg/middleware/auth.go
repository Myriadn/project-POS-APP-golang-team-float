package middleware

import (
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type SessionValidator interface {
	ValidateSession(token uuid.UUID) (*entity.User, error)
}

type AuthMiddleware struct {
	validator SessionValidator
}

func NewAuthMiddleware(validator SessionValidator) *AuthMiddleware {
	return &AuthMiddleware{validator: validator}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			utils.Unauthorized(c, "Invalid authorization format")
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

// //untuk melihat apakah di izinkan atau tidak
// func (m *AuthMiddleware) RequirePermission(code string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		sessionIDStr, err := c.Cookie("session")
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: session cookie missing"})
// 			return
// 		}

// 		userID, err := m.Service.SessionService.GetUserIDBySession(sessionID)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error or invalid session"})
// 			return
// 		}

// 		allowed, err := m.Service.Permission.Allowed(userID, code)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error checking permission"})
// 			return
// 		}
// 		if !allowed {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
// 			return
// 		}
// 		c.Next()
// 	}
// }
