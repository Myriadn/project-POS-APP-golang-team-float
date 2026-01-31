package repository

import (
	"context"
	"errors"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/dto"

	"time"

	"gorm.io/gorm"
)

type OrderRepository interface {
	ListOrders(ctx context.Context) ([]dto.OrderResponse, error)
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID uint) (dto.OrderResponse, error)
	UpdateOrder(ctx context.Context, id int, req dto.UpdateOrderRequest) (dto.OrderResponse, error)
	DeleteOrder(ctx context.Context, id int) error
	ListAvailableTables(ctx context.Context) ([]dto.TableResponse, error)
	ListPaymentMethods(ctx context.Context) ([]dto.PaymentMethodResponse, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) ListOrders(ctx context.Context) ([]dto.OrderResponse, error) {
	var orders []entity.Order
	err := r.db.WithContext(ctx).Preload("OrderItems.Product").Preload("Table").Preload("PaymentMethod").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	var result []dto.OrderResponse
	for _, o := range orders {
		var items []dto.OrderItemDetail
		for _, item := range o.OrderItems {
			items = append(items, dto.OrderItemDetail{
				ProductID:   int(item.ProductID),
				ProductName: item.Product.Name,
				Quantity:    item.Quantity,
				Price:       item.UnitPrice,
				Subtotal:    item.TotalPrice,
			})
		}
		result = append(result, dto.OrderResponse{
			ID:           int(o.ID),
			CustomerName: o.CustomerName,
			TableID:      intPtrToInt(o.TableID),
			OrderItems:   items,
			PaymentMethod: func() string {
				if o.PaymentMethod != nil {
					return o.PaymentMethod.Name
				} else {
					return ""
				}
			}(),
			Tax:   o.TaxAmount,
			Total: o.Total,
		})
	}
	return result, nil
}

func (r *orderRepository) CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID uint) (dto.OrderResponse, error) {
	const taxRate = 0.10
	order := entity.Order{
		CustomerName:    req.CustomerName,
		TableID:         uintPtr(uint(req.TableID)),
		PaymentMethodID: uintPtr(uint(req.PaymentMethodID)),
		UserID:          userID,  // ✅ Set dari parameter
		Status:          "ready",
		TaxRate:         taxRate * 100,
		OrderDate:       Now(),
	}
	subtotal := 0.0
	var items []entity.OrderItem
	// Cek dan kurangi stok produk
	tx := r.db.WithContext(ctx).Begin()
	for _, item := range req.OrderItems {
		var product entity.Product
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			return dto.OrderResponse{}, err
		}
		if product.Stock < item.Quantity {
			tx.Rollback()
			return dto.OrderResponse{}, errors.New("stok produk tidak cukup untuk produk: " + product.Name)
		}
		// Kurangi stok
		if err := tx.Model(&product).UpdateColumn("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return dto.OrderResponse{}, err
		}
		itemTotal := product.Price * float64(item.Quantity)
		subtotal += itemTotal
		items = append(items, entity.OrderItem{
			ProductID:  uint(item.ProductID),
			Quantity:   item.Quantity,
			UnitPrice:  product.Price,
			TotalPrice: itemTotal,
		})
	}
	order.Subtotal = subtotal
	order.TaxAmount = subtotal * taxRate
	order.Total = subtotal + order.TaxAmount
	order.OrderItems = items

	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return dto.OrderResponse{}, err
	}
	tx.Commit()
	return toOrderResponse(order), nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, id int, req dto.UpdateOrderRequest) (dto.OrderResponse, error) {
	var order entity.Order
	if err := r.db.WithContext(ctx).Preload("OrderItems").First(&order, id).Error; err != nil {
		return dto.OrderResponse{}, err
	}
	order.CustomerName = req.CustomerName
	order.PaymentMethodID = uintPtr(uint(req.PaymentMethodID))
	r.db.WithContext(ctx).Where("order_id = ?", order.ID).Delete(&entity.OrderItem{})
	subtotal := 0.0
	var items []entity.OrderItem
	for _, item := range req.OrderItems {
		var product entity.Product
		if err := r.db.WithContext(ctx).First(&product, item.ProductID).Error; err != nil {
			return dto.OrderResponse{}, err
		}
		itemTotal := product.Price * float64(item.Quantity)
		subtotal += itemTotal
		items = append(items, entity.OrderItem{
			OrderID:    order.ID,
			ProductID:  uint(item.ProductID),
			Quantity:   item.Quantity,
			UnitPrice:  product.Price,
			TotalPrice: itemTotal,
		})
	}
	order.Subtotal = subtotal
	order.TaxAmount = subtotal * 0.10
	order.Total = subtotal + order.TaxAmount
	order.OrderItems = items
	err := r.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	if err != nil {
		return dto.OrderResponse{}, err
	}
	return toOrderResponse(order), nil
}

func (r *orderRepository) DeleteOrder(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Order{}, id).Error
}

func (r *orderRepository) ListAvailableTables(ctx context.Context) ([]dto.TableResponse, error) {
	var tables []entity.Table
	err := r.db.WithContext(ctx).Where("status = ?", "available").Find(&tables).Error
	if err != nil {
		return nil, err
	}
	var result []dto.TableResponse
	for _, t := range tables {
		result = append(result, dto.TableResponse{
			ID:          int(t.ID),
			TableNumber: t.TableNumber,
			Available:   t.Status == "available",
		})
	}
	return result, nil
}

func (r *orderRepository) ListPaymentMethods(ctx context.Context) ([]dto.PaymentMethodResponse, error) {
	var methods []entity.PaymentMethod
	err := r.db.WithContext(ctx).Where("is_active = ?", true).Find(&methods).Error
	if err != nil {
		return nil, err
	}
	var result []dto.PaymentMethodResponse
	for _, m := range methods {
		result = append(result, dto.PaymentMethodResponse{
			ID:   int(m.ID),
			Name: m.Name,
		})
	}
	return result, nil
}

// Helper
func toOrderResponse(o entity.Order) dto.OrderResponse {
	var items []dto.OrderItemDetail
	for _, item := range o.OrderItems {
		items = append(items, dto.OrderItemDetail{
			ProductID:   int(item.ProductID),
			ProductName: item.Product.Name,
			Quantity:    item.Quantity,
			Price:       item.UnitPrice,
			Subtotal:    item.TotalPrice,
		})
	}
	return dto.OrderResponse{
		ID:           int(o.ID),
		CustomerName: o.CustomerName,
		TableID:      intPtrToInt(o.TableID),
		OrderItems:   items,
		PaymentMethod: func() string {
			if o.PaymentMethod != nil {
				return o.PaymentMethod.Name
			} else {
				return ""
			}
		}(),
		Tax:   o.TaxAmount,
		Total: o.Total,
	}
}

func uintPtr(i uint) *uint { return &i }
func intPtrToInt(i *uint) int {
	if i == nil {
		return 0
	} else {
		return int(*i)
	}
}

func Now() time.Time { return time.Now() }
