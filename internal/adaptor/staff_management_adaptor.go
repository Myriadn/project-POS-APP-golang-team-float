package adaptor

import (
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"
	"strconv"

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
	var req dto.UpdateStaffManagementReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}
	ctx := c.Request.Context()
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}
	id := uint(idInt)
	result, err := a.StaffManagementUsecase.UpdateStaffManagementUsecase(ctx, id, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

func (a *StaffManagementAdaptor) GetDetailStaffManagement(c *gin.Context) {

	ctx := c.Request.Context()
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}
	id := uint(idInt)
	user, message, err := a.StaffManagementUsecase.GetDetailStaffManagement(ctx, id)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, 200, message.Message, user)
}

func (a *StaffManagementAdaptor) DeleteStaffManagement(c *gin.Context) {
	ctx := c.Request.Context()
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}
	id := uint(idInt)
	result, err := a.StaffManagementUsecase.DeleteStaffManagementUsecase(ctx, id)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

func (a *StaffManagementAdaptor) GetAllStaffManagement(c *gin.Context) {
	ctx := c.Request.Context()
	sortBy := c.Query("sort_by")
	req := dto.GetStaffManagementFilterRequest{
		SortBy: sortBy,
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "Parameter salah", nil)
		return
	}
	result, pagination, err := a.StaffManagementUsecase.GetAllStaffManagement(ctx, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.SuccessPaginationResponse(c, 200, "Berhasil mengambil daftar staff", result, &pagination)
}
