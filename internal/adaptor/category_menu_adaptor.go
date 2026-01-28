package adaptor

import (
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"

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
