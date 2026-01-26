package dto

import "time"

// DashboardSummaryResponse represents the main dashboard summary data
type DashboardSummaryResponse struct {
	DailySales   SalesSummary  `json:"daily_sales"`
	MonthlySales SalesSummary  `json:"monthly_sales"`
	TableSummary TableSummary  `json:"table_summary"`
	PopularItems []PopularItem `json:"popular_items"`
	NewProducts  []NewProduct  `json:"new_products"`
}

// SalesSummary represents sales data for a period
type SalesSummary struct {
	TotalOrders     int64   `json:"total_orders"`
	TotalRevenue    float64 `json:"total_revenue"`
	CompletedOrders int64   `json:"completed_orders"`
	CancelledOrders int64   `json:"cancelled_orders"`
}

// TableSummary represents table availability summary
type TableSummary struct {
	TotalTables     int64 `json:"total_tables"`
	AvailableTables int64 `json:"available_tables"`
	OccupiedTables  int64 `json:"occupied_tables"`
	ReservedTables  int64 `json:"reserved_tables"`
}

// PopularItem represents a popular product with its order count
type PopularItem struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	CategoryName string  `json:"category_name"`
	TotalOrdered int64   `json:"total_ordered"`
}

// NewProduct represents a product created within the last 30 days
type NewProduct struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Price        float64   `json:"price"`
	CategoryName string    `json:"category_name"`
	Availability string    `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
}

// DailySalesRequest represents optional filter for daily sales
type DailySalesRequest struct {
	Date string `form:"date"` // format: YYYY-MM-DD, default: today
}

// MonthlySalesRequest represents optional filter for monthly sales
type MonthlySalesRequest struct {
	Month int `form:"month"` // 1-12, default: current month
	Year  int `form:"year"`  // default: current year
}
