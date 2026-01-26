package repository

import (
	"time"

	"project-POS-APP-golang-team-float/internal/data/entity"
)

// DashboardStats represents aggregated dashboard statistics
type DashboardStats struct {
	TotalOrders     int64
	TotalRevenue    float64
	CompletedOrders int64
	CancelledOrders int64
}

// PopularProduct represents a product with its order count
type PopularProduct struct {
	ID           uint
	Name         string
	Image        string
	Price        float64
	CategoryName string
	TotalOrdered int64
}

// TableStats represents table availability statistics
type TableStats struct {
	TotalTables     int64
	AvailableTables int64
	OccupiedTables  int64
	ReservedTables  int64
}

// GetDailySalesStats retrieves sales statistics for a specific date
func (r *Repository) GetDailySalesStats(date time.Time) (*DashboardStats, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var stats DashboardStats

	// Get total orders and revenue
	err := r.db.Model(&entity.Order{}).
		Where("order_date >= ? AND order_date < ?", startOfDay, endOfDay).
		Select("COUNT(*) as total_orders, COALESCE(SUM(total), 0) as total_revenue").
		Row().Scan(&stats.TotalOrders, &stats.TotalRevenue)
	if err != nil {
		return nil, err
	}

	// Get completed orders count
	err = r.db.Model(&entity.Order{}).
		Where("order_date >= ? AND order_date < ? AND status = ?", startOfDay, endOfDay, "completed").
		Count(&stats.CompletedOrders).Error
	if err != nil {
		return nil, err
	}

	// Get cancelled orders count
	err = r.db.Model(&entity.Order{}).
		Where("order_date >= ? AND order_date < ? AND status = ?", startOfDay, endOfDay, "cancelled").
		Count(&stats.CancelledOrders).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetMonthlySalesStats retrieves sales statistics for a specific month
func (r *Repository) GetMonthlySalesStats(year, month int) (*DashboardStats, error) {
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	var stats DashboardStats

	// Get total orders and revenue
	err := r.db.Model(&entity.Order{}).
		Where("order_date >= ? AND order_date < ?", startOfMonth, endOfMonth).
		Select("COUNT(*) as total_orders, COALESCE(SUM(total), 0) as total_revenue").
		Row().Scan(&stats.TotalOrders, &stats.TotalRevenue)
	if err != nil {
		return nil, err
	}

	// Get completed orders count
	err = r.db.Model(&entity.Order{}).
		Where("order_date >= ? AND order_date < ? AND status = ?", startOfMonth, endOfMonth, "completed").
		Count(&stats.CompletedOrders).Error
	if err != nil {
		return nil, err
	}

	// Get cancelled orders count
	err = r.db.Model(&entity.Order{}).
		Where("order_date >= ? AND order_date < ? AND status = ?", startOfMonth, endOfMonth, "cancelled").
		Count(&stats.CancelledOrders).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetTableStats retrieves table availability statistics
func (r *Repository) GetTableStats() (*TableStats, error) {
	var stats TableStats

	// Get total tables
	err := r.db.Model(&entity.Table{}).Count(&stats.TotalTables).Error
	if err != nil {
		return nil, err
	}

	// Get available tables
	err = r.db.Model(&entity.Table{}).
		Where("status = ?", "available").
		Count(&stats.AvailableTables).Error
	if err != nil {
		return nil, err
	}

	// Get occupied tables
	err = r.db.Model(&entity.Table{}).
		Where("status = ?", "occupied").
		Count(&stats.OccupiedTables).Error
	if err != nil {
		return nil, err
	}

	// Get reserved tables
	err = r.db.Model(&entity.Table{}).
		Where("status = ?", "reserved").
		Count(&stats.ReservedTables).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetPopularProducts retrieves the most ordered products
func (r *Repository) GetPopularProducts(limit int) ([]PopularProduct, error) {
	var products []PopularProduct

	err := r.db.Table("order_items").
		Select(`
			products.id,
			products.name,
			products.image,
			products.price,
			categories.name as category_name,
			SUM(order_items.quantity) as total_ordered
		`).
		Joins("JOIN products ON products.id = order_items.product_id").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.deleted_at IS NULL").
		Group("products.id, products.name, products.image, products.price, categories.name").
		Order("total_ordered DESC").
		Limit(limit).
		Scan(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetNewProducts retrieves products created within the last 30 days
func (r *Repository) GetNewProducts(limit int) ([]entity.Product, error) {
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	var products []entity.Product
	err := r.db.Preload("Category").
		Where("created_at >= ?", thirtyDaysAgo).
		Order("created_at DESC").
		Limit(limit).
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}
