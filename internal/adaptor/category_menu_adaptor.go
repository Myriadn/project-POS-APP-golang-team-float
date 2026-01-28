package adaptor

import (
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryMenuAdaptor struct {
	CategoryMenuUsecase usecase.CategoryMenuUsecaseInterface
}

func NewCategoryMenuAdaptor(uc usecase.CategoryMenuUsecaseInterface) *CategoryMenuAdaptor {
	return &CategoryMenuAdaptor{CategoryMenuUsecase: uc}
}

// create new category dari request di user
func (a *CategoryMenuAdaptor) CreateNewCategoryMenu(c *gin.Context) {
	var req dto.CreateNewCategoryMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}
	ctx := c.Request.Context()

	result, err := a.CategoryMenuUsecase.CreateNewCategoryUsecase(ctx, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

// request update category menu
func (a *CategoryMenuAdaptor) UpdateCategoryMenu(c *gin.Context) {
	var req dto.UpdateCategoryMenuReq
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
	result, err := a.CategoryMenuUsecase.UpdateCategoryMenuUsecase(ctx, id, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

// requset untuk dapat detail category  menu
func (a *CategoryMenuAdaptor) GetDetailCategoryMenu(c *gin.Context) {

	ctx := c.Request.Context()
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}
	id := uint(idInt)
	user, message, err := a.CategoryMenuUsecase.GetDetailCategoryMenu(ctx, id)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, 200, message.Message, user)
}

// mendapatkan  semua  category menu
func (a *CategoryMenuAdaptor) GetAllCategoryMenu(c *gin.Context) {
	ctx := c.Request.Context()
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := dto.FilterRequest{
		Page:  page,
		Limit: limit,
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "Parameter salah", nil)
		return
	}
	result, pagination, err := a.CategoryMenuUsecase.GetAllCategoryMenu(ctx, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.SuccessPaginationResponse(c, 200, "Berhasil mengambil semua daftar category menu", result, &pagination)
}

// request untuk delete category menu
func (a *CategoryMenuAdaptor) DeleteCategoryMenu(c *gin.Context) {
	ctx := c.Request.Context()
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}
	id := uint(idInt)
	result, err := a.CategoryMenuUsecase.DeleteCategoryMenu(ctx, id)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}
