package dto

type OrderItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CreateOrderRequest struct {
	CustomerName    string             `json:"customer_name"`
	TableID         int                `json:"table_id"`
	OrderItems      []OrderItemRequest `json:"order_items"`
	PaymentMethodID int                `json:"payment_method_id"`
}

type UpdateOrderRequest struct {
	CustomerName    string             `json:"customer_name"`
	OrderItems      []OrderItemRequest `json:"order_items"`
	PaymentMethodID int                `json:"payment_method_id"`
}

type OrderResponse struct {
	ID            int               `json:"id"`
	CustomerName  string            `json:"customer_name"`
	TableID       int               `json:"table_id"`
	OrderItems    []OrderItemDetail `json:"order_items"`
	PaymentMethod string            `json:"payment_method"`
	Tax           float64           `json:"tax"`
	Total         float64           `json:"total"`
}

type OrderItemDetail struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Subtotal    float64 `json:"subtotal"`
}

type TableResponse struct {
	ID          int    `json:"id"`
	TableNumber string `json:"table_number"`
	Available   bool   `json:"available"`
}

type PaymentMethodResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
