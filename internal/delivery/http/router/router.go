package router

import (
	"github.com/gin-gonic/gin"

	"project-POS-APP-golang-team-float/internal/delivery/http/handler"
	"project-POS-APP-golang-team-float/internal/delivery/http/middleware"
)

type Router struct {
	engine         *gin.Engine
	authHandler    *handler.AuthHandler
	authMiddleware *middleware.AuthMiddleware
}

func NewRouter(
	authHandler *handler.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
) *Router {
	return &Router{
		engine:         gin.Default(),
		authHandler:    authHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) Setup() *gin.Engine {
	api := r.engine.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", r.authHandler.Login)
			auth.POST("/verify-otp", r.authHandler.VerifyOTP)
			auth.POST("/check-email", r.authHandler.CheckEmail)
			auth.POST("/reset-password", r.authHandler.ResetPassword)
			auth.POST("/logout", r.authMiddleware.Authenticate(), r.authHandler.Logout)
		}
	}

	return r.engine
}
