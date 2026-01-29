package adaptor

import (
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProfileAdaptor struct {
	ProfileUsecase usecase.ProfileUsecaseInterface
}

func NewProfileAdaptor(uc usecase.ProfileUsecaseInterface) *ProfileAdaptor {
	return &ProfileAdaptor{ProfileUsecase: uc}
}

// request update profile menu
func (a *ProfileAdaptor) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}
	ctx := c.Request.Context()
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

	result, err := a.ProfileUsecase.UpdateProfileUsecase(ctx, userID, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}
