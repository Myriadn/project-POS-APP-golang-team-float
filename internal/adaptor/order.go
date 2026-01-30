package adaptor

import (
	"net/http"
	"strconv"

	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OrderAdaptor struct {
	orderUC usecase.OrderUsecaseInterface
}

func NewOrderAdaptor(orderUC usecase.OrderUsecaseInterface) *OrderAdaptor {
	return &OrderAdaptor{orderUC: orderUC}
}

// List Orders
func (a *OrderAdaptor) ListOrders(c *gin.Context) {
	orders, err := a.orderUC.ListOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// Create Order
func (a *OrderAdaptor) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := a.orderUC.CreateOrder(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// Update Order
func (a *OrderAdaptor) UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := a.orderUC.UpdateOrder(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Delete Order
func (a *OrderAdaptor) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := a.orderUC.DeleteOrder(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}

// List Tables (only available)
func (a *OrderAdaptor) ListAvailableTables(c *gin.Context) {
	tables, err := a.orderUC.ListAvailableTables(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tables})
}

// List Payment Methods
func (a *OrderAdaptor) ListPaymentMethods(c *gin.Context) {
	methods, err := a.orderUC.ListPaymentMethods(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": methods})
}
