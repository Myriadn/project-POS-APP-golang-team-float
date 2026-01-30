package adaptor

import (
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductMenuAdaptor struct {
	ProductMenuUsecase usecase.ProductMenuUsecaseInterface
}

func NewProductMenuAdaptor(uc usecase.ProductMenuUsecaseInterface) *ProductMenuAdaptor {
	return &ProductMenuAdaptor{ProductMenuUsecase: uc}
}

// create new product dari request di user
func (a *ProductMenuAdaptor) CreateNewProductMenu(c *gin.Context) {
	var req dto.CreateNewProductMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err.Error())
		return
	}
	ctx := c.Request.Context()

	result, err := a.ProductMenuUsecase.CreateNewProductUsecase(ctx, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

// request update product menu
func (a *ProductMenuAdaptor) UpdateProductMenu(c *gin.Context) {
	var req dto.UpdateProductMenuReq
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
	result, err := a.ProductMenuUsecase.UpdateProductMenuUsecase(ctx, id, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}

// requset untuk dapat detail product  menu
func (a *ProductMenuAdaptor) GetDetailProductMenu(c *gin.Context) {

	ctx := c.Request.Context()
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}
	id := uint(idInt)
	user, message, err := a.ProductMenuUsecase.GetDetailProductMenu(ctx, id)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, 200, message.Message, user)
}

func (a *ProductMenuAdaptor) GetAllStaffProductMenu(c *gin.Context) {
	ctx := c.Request.Context()
	MenuType := c.Query("menu_type")
	req := dto.FilterRequest{
		MenuType: MenuType,
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "Parameter salah", nil)
		return
	}
	result, pagination, err := a.ProductMenuUsecase.GetAllProductMenu(ctx, req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.SuccessPaginationResponse(c, 200, "Berhasil mengambil daftar product", result, &pagination)
}

// request untuk delete category menu
func (a *ProductMenuAdaptor) DeleteProductMenu(c *gin.Context) {
	ctx := c.Request.Context()
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}
	id := uint(idInt)
	result, err := a.ProductMenuUsecase.DeleteProductMenu(ctx, id)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	utils.Success(c, result.Message, nil)
}
