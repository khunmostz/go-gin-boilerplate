package handler

import (
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"go-gin-boilerplate/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BarHandler struct {
	barSvc port.BarService
}

func NewBarHandler(barSvc port.BarService) *BarHandler {
	return &BarHandler{barSvc: barSvc}
}

func (bh *BarHandler) CreateBar(c *gin.Context) {
	var bar domain.Bar
	if err := c.ShouldBindJSON(&bar); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	createdBar, err := bh.barSvc.CreateBar(c.Request.Context(), &bar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponseWithMessage("Bar created successfully", createdBar))
}

func (bh *BarHandler) GetBars(c *gin.Context) {
	bars, err := bh.barSvc.GetBars(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponseWithMessage("Bars retrieved successfully", bars))
}

func (bh *BarHandler) GetBarByID(c *gin.Context) {
	id := c.Param("id")
	bar, err := bh.barSvc.GetBarByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponseWithMessage("Bar retrieved successfully", bar))
}

func (bh *BarHandler) UpdateBar(c *gin.Context) {
	id := c.Param("id")
	var bar domain.Bar
	if err := c.ShouldBindJSON(&bar); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	updatedBar, err := bh.barSvc.UpdateBar(c.Request.Context(), id, &bar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponseWithMessage("Bar updated successfully", updatedBar))
}

func (bh *BarHandler) DeleteBar(c *gin.Context) {
	id := c.Param("id")
	if err := bh.barSvc.DeleteBar(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponseWithMessage("Bar deleted successfully", ""))
}
