package adaptor

import (
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"

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
