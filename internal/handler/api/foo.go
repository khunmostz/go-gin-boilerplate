package api

import (
	"go-gin-boilerplate/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterFooRoutes(router *gin.RouterGroup, fooHandler *handler.FooHandler) {
	router.POST("/", fooHandler.CreateFoo)
}
