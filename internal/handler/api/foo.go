package api

import (
	"go-gin-boilerplate/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterFooRoutes(router *gin.RouterGroup, fooHandler *handler.FooHandler) {
	router.POST("/", fooHandler.Create)
	router.GET("/", fooHandler.GetAll)
	router.GET("/:id", fooHandler.GetByID)
	router.PUT("/:id", fooHandler.UpdateByID)
	router.DELETE("/:id", fooHandler.DeleteByID)
}
