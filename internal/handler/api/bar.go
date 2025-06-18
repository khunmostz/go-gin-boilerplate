package api

import (
	"go-gin-boilerplate/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBarRoutes(router *gin.RouterGroup, barHandler *handler.BarHandler) {
	router.POST("/", barHandler.CreateBar)
	router.GET("/", barHandler.GetBars)
	router.GET("/:id", barHandler.GetBarByID)
	router.PUT("/:id", barHandler.UpdateBar)
	router.DELETE("/:id", barHandler.DeleteBar)
}
