package dto

type RevenueByStatusResponse struct {
	Status       string  `json:"status"`
	TotalOrders  int64   `json:"total_orders"`
	TotalRevenue float64 `json:"total_revenue"`
}

type MonthlyRevenueResponse struct {
	Month        string  `json:"month"` // contoh "January", "February"
	MonthInt     int     `json:"month_int"`
	TotalRevenue float64 `json:"total_revenue"`
}

type ProductRevenueResponse struct {
	ProductID    uint    `json:"product_id"`
	ProductName  string  `json:"product_name"`
	CategoryName string  `json:"category_name"`
	TotalSold    int64   `json:"total_sold"`
	TotalRevenue float64 `json:"total_revenue"`
}
