package routers

import (
	"simple-live-attendance/src/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	handler := handlers.NewHandler(db)

	router.POST("/clock-in", handler.ClockIn)
	router.POST("/clock-out", handler.ClockOut)
}
