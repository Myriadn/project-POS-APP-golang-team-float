package adaptor

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type DashboardAdaptor struct {
	uc *usecase.Usecase
}

func NewDashboardAdaptor(uc *usecase.Usecase) *DashboardAdaptor {
	return &DashboardAdaptor{uc: uc}
}

// Get dashboard summary
// Get all dashboard data including daily/monthly sales, table summary, popular items, and new products
// /dashboard [get]
func (a *DashboardAdaptor) GetDashboardSummary(c *gin.Context) {
	summary, err := a.uc.GetDashboardSummary()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get dashboard summary", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboard summary retrieved successfully", summary)
}

// Get daily sales statistics
// Get sales statistics for a specific date (default: today)
// /dashboard/daily-sales [get]
func (a *DashboardAdaptor) GetDailySales(c *gin.Context) {
	var req dto.DailySalesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	sales, err := a.uc.GetDailySales(req.Date)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get daily sales", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Daily sales retrieved successfully", sales)
}

// Get monthly sales statistics
// Get sales statistics for a specific month (default: current month)
// /dashboard/monthly-sales [get]
func (a *DashboardAdaptor) GetMonthlySales(c *gin.Context) {
	var req dto.MonthlySalesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	sales, err := a.uc.GetMonthlySales(req.Year, req.Month)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get monthly sales", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Monthly sales retrieved successfully", sales)
}

// Get table summary
// Get table availability statistics
// /dashboard/tables [get]
func (a *DashboardAdaptor) GetTableSummary(c *gin.Context) {
	summary, err := a.uc.GetTableSummary()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get table summary", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Table summary retrieved successfully", summary)
}

// Get popular products
// Get the most ordered products
// /dashboard/popular-products [get]
func (a *DashboardAdaptor) GetPopularProducts(c *gin.Context) {
	limit := 10
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	products, err := a.uc.GetPopularProducts(limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get popular products", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Popular products retrieved successfully", products)
}

// Get new products
// Get products created within the last 30 days
// /dashboard/new-products [get]
func (a *DashboardAdaptor) GetNewProducts(c *gin.Context) {
	limit := 10
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	products, err := a.uc.GetNewProducts(limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get new products", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "New products retrieved successfully", products)
}
