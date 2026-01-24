package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/response"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req usecase.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.authUsecase.Login(req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, result.Message, nil)
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req usecase.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	result, err := h.authUsecase.VerifyOTP(req, ipAddress, userAgent)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "Login successful", result)
}

func (h *AuthHandler) CheckEmail(c *gin.Context) {
	var req usecase.CheckEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.authUsecase.CheckEmail(req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, result.Message, nil)
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req usecase.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	err := h.authUsecase.ResetPassword(req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "Password reset successful", nil)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}

	token, err := uuid.Parse(tokenStr)
	if err != nil {
		response.BadRequest(c, "Invalid token", nil)
		return
	}

	if err := h.authUsecase.Logout(token); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "Logged out successfully", nil)
}
