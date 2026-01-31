package entity

import (
	"time"
)

type Order struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	OrderNumber     string         `gorm:"uniqueIndex;size:20;not null" json:"order_number"`
	TableID         *uint          `json:"table_id"`
	UserID          uint           `gorm:"not null" json:"user_id"`
	PaymentMethodID *uint          `json:"payment_method_id"`
	CustomerName    string         `gorm:"size:255;not null" json:"customer_name"`
	Status          string         `gorm:"size:30;default:'ready'" json:"status"` // ready, in_process, completed, cancelled, cooking_now, in_the_kitchen, ready_to_serve
	Subtotal        float64        `gorm:"type:decimal(15,2);not null;default:0" json:"subtotal"`
	TaxRate         float64        `gorm:"type:decimal(5,2);default:5.00" json:"tax_rate"`
	TaxAmount       float64        `gorm:"type:decimal(15,2);default:0" json:"tax_amount"`
	Total           float64        `gorm:"type:decimal(15,2);not null;default:0" json:"total"`
	Notes           string         `gorm:"type:text" json:"notes"`
	OrderDate       time.Time      `gorm:"default:now()" json:"order_date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Table           *Table         `gorm:"foreignKey:TableID" json:"table,omitempty"`
	User            User           `gorm:"foreignKey:UserID" json:"user,omitzero"`
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"payment_method,omitempty"`
	OrderItems      []OrderItem    `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order_items,omitempty"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	OrderID    uint      `gorm:"not null" json:"order_id"`
	ProductID  uint      `gorm:"not null" json:"product_id"`
	Quantity   int       `gorm:"not null;default:1" json:"quantity"`
	UnitPrice  float64   `gorm:"type:decimal(15,2);not null" json:"unit_price"`
	TotalPrice float64   `gorm:"type:decimal(15,2);not null" json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	Order      Order     `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"-"`
	Product    Product   `gorm:"foreignKey:ProductID" json:"product,omitzero"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
