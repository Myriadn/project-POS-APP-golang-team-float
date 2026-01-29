package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"time"
)

type ReportUsecase struct {
	repo repository.ReportRepoInterface
}

type ReportUsecaseInterface interface {
	GetRevenueReport(ctx context.Context, year int) (map[string]interface{}, error)
}

func NewReportUsecase(repo repository.ReportRepoInterface) ReportUsecaseInterface {
	return &ReportUsecase{repo: repo}
}

func (u *ReportUsecase) GetRevenueReport(ctx context.Context, year int) (map[string]interface{}, error) {
	// Default year to current year if 0
	if year == 0 {
		year = time.Now().Year()
	}

	// Get Revenue By Status
	byStatus, err := u.repo.GetRevenueByStatus(ctx)
	if err != nil {
		return nil, err
	}

	// Get Monthly Revenue
	monthly, err := u.repo.GetMonthlyRevenue(ctx, year)
	if err != nil {
		return nil, err
	}

	// Convert MonthInt to String Name (1 -> January)
	for i, m := range monthly {
		monthly[i].Month = time.Month(m.MonthInt).String()
	}

	// Get Product Revenue (Top 20 products for example)
	products, err := u.repo.GetProductRevenue(ctx, 20)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"year":                year,
		"revenue_by_status":   byStatus,
		"monthly_revenue":     monthly,
		"top_product_revenue": products,
	}, nil
}
