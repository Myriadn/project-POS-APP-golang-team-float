package adaptor

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type AuthAdaptor struct {
	usecase *usecase.Usecase
}

func NewAuthAdaptor(uc *usecase.Usecase) *AuthAdaptor {
	return &AuthAdaptor{usecase: uc}
}

func (a *AuthAdaptor) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	result, err := a.usecase.Login(req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

func (a *AuthAdaptor) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	result, err := a.usecase.VerifyOTP(req, ipAddress, userAgent)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, "Login successful", result)
}

func (a *AuthAdaptor) CheckEmail(c *gin.Context) {
	var req dto.CheckEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	result, err := a.usecase.CheckEmail(req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

func (a *AuthAdaptor) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}

	err := a.usecase.ResetPassword(req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, "Password reset successful", nil)
}

func (a *AuthAdaptor) Logout(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}

	token, err := uuid.Parse(tokenStr)
	if err != nil {
		utils.BadRequest(c, "Invalid token", nil)
		return
	}

	if err := a.usecase.Logout(token); err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, "Logged out successfully", nil)
}
