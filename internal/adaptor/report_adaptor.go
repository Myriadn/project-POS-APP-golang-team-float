package adaptor

import (
	"net/http"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportAdaptor struct {
	usecase usecase.ReportUsecaseInterface
}

func NewReportAdaptor(uc usecase.ReportUsecaseInterface) *ReportAdaptor {
	return &ReportAdaptor{usecase: uc}
}

func (a *ReportAdaptor) GetRevenueReport(c *gin.Context) {
	ctx := c.Request.Context()

	// Get year from query param, default to 0 (handled in usecase)
	yearStr := c.Query("year")
	year, _ := strconv.Atoi(yearStr)

	result, err := a.usecase.GetRevenueReport(ctx, year)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get revenue report", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Revenue report retrieved successfully", result)
}
