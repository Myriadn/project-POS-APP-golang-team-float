package usecase

import (
	"time"

	"project-POS-APP-golang-team-float/internal/dto"
)

// GetDashboardSummary retrieves all dashboard data including sales, tables, and products
func (u *Usecase) GetDashboardSummary() (*dto.DashboardSummaryResponse, error) {
	// Get daily sales (today)
	dailyStats, err := u.repo.GetDailySalesStats(time.Now())
	if err != nil {
		return nil, err
	}

	// Get monthly sales (current month)
	now := time.Now()
	monthlyStats, err := u.repo.GetMonthlySalesStats(now.Year(), int(now.Month()))
	if err != nil {
		return nil, err
	}

	// Get table statistics
	tableStats, err := u.repo.GetTableStats()
	if err != nil {
		return nil, err
	}

	// Get popular products (top 10)
	popularProducts, err := u.repo.GetPopularProducts(10)
	if err != nil {
		return nil, err
	}

	// Get new products (last 30 days, limit 10)
	newProducts, err := u.repo.GetNewProducts(10)
	if err != nil {
		return nil, err
	}

	// Build response
	response := &dto.DashboardSummaryResponse{
		DailySales: dto.SalesSummary{
			TotalOrders:     dailyStats.TotalOrders,
			TotalRevenue:    dailyStats.TotalRevenue,
			CompletedOrders: dailyStats.CompletedOrders,
			CancelledOrders: dailyStats.CancelledOrders,
		},
		MonthlySales: dto.SalesSummary{
			TotalOrders:     monthlyStats.TotalOrders,
			TotalRevenue:    monthlyStats.TotalRevenue,
			CompletedOrders: monthlyStats.CompletedOrders,
			CancelledOrders: monthlyStats.CancelledOrders,
		},
		TableSummary: dto.TableSummary{
			TotalTables:     tableStats.TotalTables,
			AvailableTables: tableStats.AvailableTables,
			OccupiedTables:  tableStats.OccupiedTables,
			ReservedTables:  tableStats.ReservedTables,
		},
		PopularItems: make([]dto.PopularItem, 0, len(popularProducts)),
		NewProducts:  make([]dto.NewProduct, 0, len(newProducts)),
	}

	// Map popular products to response
	for _, p := range popularProducts {
		response.PopularItems = append(response.PopularItems, dto.PopularItem{
			ID:           p.ID,
			Name:         p.Name,
			Image:        p.Image,
			Price:        p.Price,
			CategoryName: p.CategoryName,
			TotalOrdered: p.TotalOrdered,
		})
	}

	// Map new products to response
	for _, p := range newProducts {
		categoryName := ""
		if p.Category.Name != "" {
			categoryName = p.Category.Name
		}
		response.NewProducts = append(response.NewProducts, dto.NewProduct{
			ID:           p.ID,
			Name:         p.Name,
			Image:        p.Image,
			Price:        p.Price,
			CategoryName: categoryName,
			Availability: p.Availability,
			CreatedAt:    p.CreatedAt,
		})
	}

	return response, nil
}

// GetDailySales retrieves sales statistics for a specific date
func (u *Usecase) GetDailySales(dateStr string) (*dto.SalesSummary, error) {
	var date time.Time
	var err error

	if dateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, err
		}
	}

	stats, err := u.repo.GetDailySalesStats(date)
	if err != nil {
		return nil, err
	}

	return &dto.SalesSummary{
		TotalOrders:     stats.TotalOrders,
		TotalRevenue:    stats.TotalRevenue,
		CompletedOrders: stats.CompletedOrders,
		CancelledOrders: stats.CancelledOrders,
	}, nil
}

// GetMonthlySales retrieves sales statistics for a specific month
func (u *Usecase) GetMonthlySales(year, month int) (*dto.SalesSummary, error) {
	now := time.Now()

	if year == 0 {
		year = now.Year()
	}
	if month == 0 {
		month = int(now.Month())
	}

	stats, err := u.repo.GetMonthlySalesStats(year, month)
	if err != nil {
		return nil, err
	}

	return &dto.SalesSummary{
		TotalOrders:     stats.TotalOrders,
		TotalRevenue:    stats.TotalRevenue,
		CompletedOrders: stats.CompletedOrders,
		CancelledOrders: stats.CancelledOrders,
	}, nil
}

// GetTableSummary retrieves table availability statistics
func (u *Usecase) GetTableSummary() (*dto.TableSummary, error) {
	stats, err := u.repo.GetTableStats()
	if err != nil {
		return nil, err
	}

	return &dto.TableSummary{
		TotalTables:     stats.TotalTables,
		AvailableTables: stats.AvailableTables,
		OccupiedTables:  stats.OccupiedTables,
		ReservedTables:  stats.ReservedTables,
	}, nil
}

// GetPopularProducts retrieves the most ordered products
func (u *Usecase) GetPopularProducts(limit int) ([]dto.PopularItem, error) {
	if limit <= 0 {
		limit = 10
	}

	products, err := u.repo.GetPopularProducts(limit)
	if err != nil {
		return nil, err
	}

	result := make([]dto.PopularItem, 0, len(products))
	for _, p := range products {
		result = append(result, dto.PopularItem{
			ID:           p.ID,
			Name:         p.Name,
			Image:        p.Image,
			Price:        p.Price,
			CategoryName: p.CategoryName,
			TotalOrdered: p.TotalOrdered,
		})
	}

	return result, nil
}

// GetNewProducts retrieves products created within the last 30 days
func (u *Usecase) GetNewProducts(limit int) ([]dto.NewProduct, error) {
	if limit <= 0 {
		limit = 10
	}

	products, err := u.repo.GetNewProducts(limit)
	if err != nil {
		return nil, err
	}

	result := make([]dto.NewProduct, 0, len(products))
	for _, p := range products {
		categoryName := ""
		if p.Category.Name != "" {
			categoryName = p.Category.Name
		}
		result = append(result, dto.NewProduct{
			ID:           p.ID,
			Name:         p.Name,
			Image:        p.Image,
			Price:        p.Price,
			CategoryName: categoryName,
			Availability: p.Availability,
			CreatedAt:    p.CreatedAt,
		})
	}

	return result, nil
}
