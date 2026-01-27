package adaptor

import (
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"

	"github.com/gin-gonic/gin"
)

type StaffManagementAdaptor struct {
	StaffManagementUsecase usecase.StaffManagementUsecaseInterface
}

func NewStaffManagementAdaptor(uc usecase.StaffManagementUsecaseInterface) *StaffManagementAdaptor {
	return &StaffManagementAdaptor{StaffManagementUsecase: uc}
}

func (a *StaffManagementAdaptor) CreateNewStaffManagement(c *gin.Context) {
	var req dto.CreateNewStaffManagementReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}
	ctx := c.Request.Context()

	result, err := a.StaffManagementUsecase.CreateNewStaffManagementUsecase(ctx, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

func (a *StaffManagementAdaptor) UpdateStaffManagement(c *gin.Context) {
	var req dto.CreateNewStaffManagementReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}
	ctx := c.Request.Context()
	userID := c.GetUint("user_id")
	result, err := a.StaffManagementUsecase.UpdateStaffManagementUsecase(ctx, userID, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}
