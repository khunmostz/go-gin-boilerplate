package api

import (
	"go-gin-boilerplate/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBarRoutes(router *gin.RouterGroup, barHandler *handler.BarHandler) {
	router.POST("/", barHandler.Create)
	router.GET("/", barHandler.GetAll)
	router.GET("/:id", barHandler.GetByID)
	router.PUT("/:id", barHandler.UpdateByID)
	router.DELETE("/:id", barHandler.DeleteByID)
}
