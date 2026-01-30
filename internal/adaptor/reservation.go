package adaptor

import (
	"net/http"
	"strconv"

	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ReservationAdaptor struct {
	uc usecase.ReservationUsecaseInterface
}

func NewReservationAdaptor(uc usecase.ReservationUsecaseInterface) *ReservationAdaptor {
	return &ReservationAdaptor{uc: uc}
}

func (a *ReservationAdaptor) ListReservations(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.uc.ListReservations(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (a *ReservationAdaptor) GetReservationByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	data, err := a.uc.GetReservationByID(ctx, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (a *ReservationAdaptor) CreateReservation(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := a.uc.CreateReservation(ctx, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "reservation created"})
}

func (a *ReservationAdaptor) UpdateReservation(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := a.uc.UpdateReservation(ctx, uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "reservation updated"})
}
