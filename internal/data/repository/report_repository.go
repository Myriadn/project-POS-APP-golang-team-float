package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/dto"

	"gorm.io/gorm"
)

type ReportRepo struct {
	db *gorm.DB
}

type ReportRepoInterface interface {
	GetRevenueByStatus(ctx context.Context) ([]dto.RevenueByStatusResponse, error)
	GetMonthlyRevenue(ctx context.Context, year int) ([]dto.MonthlyRevenueResponse, error)
	GetProductRevenue(ctx context.Context, limit int) ([]dto.ProductRevenueResponse, error)
}

func NewReportRepo(db *gorm.DB) ReportRepoInterface {
	return &ReportRepo{db: db}
}

// Total revenue dan breakdown berdasarkan status
func (r *ReportRepo) GetRevenueByStatus(ctx context.Context) ([]dto.RevenueByStatusResponse, error) {
	var results []dto.RevenueByStatusResponse

	err := r.db.WithContext(ctx).Table("orders").
		Select("status, COUNT(*) as total_orders, COALESCE(SUM(total), 0) as total_revenue").
		Group("status").
		Scan(&results).Error

	return results, err
}

// Total revenue per bulan (untuk satu tahun tertentu)
func (r *ReportRepo) GetMonthlyRevenue(ctx context.Context, year int) ([]dto.MonthlyRevenueResponse, error) {
	var results []dto.MonthlyRevenueResponse

	err := r.db.WithContext(ctx).Table("orders").
		Select("EXTRACT(MONTH FROM order_date) as month_int, COALESCE(SUM(total), 0) as total_revenue").
		Where("EXTRACT(YEAR FROM order_date) = ? AND status = ?", year, "completed").
		Group("EXTRACT(MONTH FROM order_date)").
		Order("month_int ASC").
		Scan(&results).Error

	return results, err
}

// List produk beserta detail revenue
func (r *ReportRepo) GetProductRevenue(ctx context.Context, limit int) ([]dto.ProductRevenueResponse, error) {
	var results []dto.ProductRevenueResponse

	query := r.db.WithContext(ctx).Table("order_items").
		Select("products.id as product_id, products.name as product_name, categories.name as category_name, SUM(order_items.quantity) as total_sold, SUM(order_items.total_price) as total_revenue").
		Joins("JOIN products ON products.id = order_items.product_id").
		Joins("JOIN categories ON categories.id = products.category_id").
		Group("products.id, products.name, categories.name").
		Order("total_revenue DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Scan(&results).Error
	return results, err
}
