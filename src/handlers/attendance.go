package handlers

import (
	"net/http"
	"simple-live-attendance/src/models"
	"simple-live-attendance/src/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	usecase *usecase.AttendanceUsecase
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		usecase: usecase.NewUsecase(db),
	}
}

func (h *Handler) ClockIn(c *gin.Context) {
	var req models.AttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.ClockIn(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clock-in successful"})
}

func (h *Handler) ClockOut(c *gin.Context) {
	var req models.AttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.ClockOut(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clock-out successful"})
}
