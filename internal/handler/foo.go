package handler

import (
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"go-gin-boilerplate/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FooHandler struct {
	fooSvc port.FooService
}

func NewFooHandler(fooSvc port.FooService) *FooHandler {
	return &FooHandler{fooSvc: fooSvc}
}

// CreateFoo creates a new foo item
// @Summary      Create a new foo
// @Description  Create a new foo item with provided information
// @Tags         foo
// @Accept       json
// @Produce      json
// @Param        foo body domain.Foo true "Foo creation data"
// @Success      200 {object} utils.Response "Successfully created foo"
// @Failure      400 {object} utils.Response "Bad request"
// @Failure      500 {object} utils.Response "Internal server error"
// @Router       /foo [post]
func (fh *FooHandler) CreateFoo(c *gin.Context) {
	foo := &domain.Foo{}
	if err := c.ShouldBindJSON(foo); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	createdFoo, err := fh.fooSvc.CreateFoo(c.Request.Context(), foo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponseWithMessage("Foo created successfully", createdFoo))
}
