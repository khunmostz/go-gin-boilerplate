package handler

import (
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BarHandler struct {
	barSvc port.BarService
}

func NewBarHandler(barSvc port.BarService) *BarHandler {
	return &BarHandler{barSvc: barSvc}
}

// Create creates a new bar item
// @Summary      Create a new bar
// @Description  Create a new bar item with provided information. Requires authentication.
// @Tags         bar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        bar body domain.Bar true "Bar creation data"
// @Success      201 {object} domain.Response{data=domain.Bar} "Successfully created bar"
// @Failure      400 {object} domain.Response "Bad request - Invalid input data"
// @Failure      401 {object} domain.Response "Unauthorized - Missing or invalid token"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /bar [post]
func (bh *BarHandler) Create(c *gin.Context) {
	var bar domain.Bar
	if err := c.ShouldBindJSON(&bar); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse(err.Error()))
		return
	}

	createdBar, err := bh.barSvc.Create(c.Request.Context(), &bar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponseWithMessage("Bar created successfully", createdBar))
}

// GetAll retrieves all bar items
// @Summary      Get all bars
// @Description  Retrieve all bar items from the database. Requires authentication.
// @Tags         bar
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} domain.Response{data=[]domain.Bar} "Successfully retrieved all bars"
// @Failure      401 {object} domain.Response "Unauthorized - Missing or invalid token"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /bar [get]
func (bh *BarHandler) GetAll(c *gin.Context) {
	bars, err := bh.barSvc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Bars retrieved successfully", bars))
}

// GetByID retrieves a bar by ID
// @Summary      Get bar by ID
// @Description  Retrieve a specific bar item by its unique identifier. Requires authentication.
// @Tags         bar
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Bar ID" example("507f1f77bcf86cd799439011")
// @Success      200 {object} domain.Response{data=domain.Bar} "Successfully retrieved bar"
// @Failure      400 {object} domain.Response "Bad request - Invalid ID format"
// @Failure      401 {object} domain.Response "Unauthorized - Missing or invalid token"
// @Failure      404 {object} domain.Response "Bar not found"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /bar/{id} [get]
func (bh *BarHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	bar, err := bh.barSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Bar retrieved successfully", bar))
}

// UpdateByID updates a bar by ID
// @Summary      Update bar by ID
// @Description  Update a specific bar item by its unique identifier. Requires authentication.
// @Tags         bar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Bar ID" example("507f1f77bcf86cd799439011")
// @Param        update body map[string]interface{} true "Update data" example({"name": "Updated Bar Name", "description": "Updated description", "status": "inactive"})
// @Success      200 {object} domain.Response{data=domain.Bar} "Successfully updated bar"
// @Failure      400 {object} domain.Response "Bad request - Invalid input data or ID format"
// @Failure      401 {object} domain.Response "Unauthorized - Missing or invalid token"
// @Failure      404 {object} domain.Response "Bar not found"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /bar/{id} [put]
func (bh *BarHandler) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var update map[string]any
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse(err.Error()))
		return
	}

	updatedBar, err := bh.barSvc.UpdateById(c.Request.Context(), id, update)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "bar not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Bar updated successfully", updatedBar))
}

// DeleteByID deletes a bar by ID
// @Summary      Delete bar by ID
// @Description  Delete a specific bar item by its unique identifier. Requires authentication.
// @Tags         bar
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Bar ID" example("507f1f77bcf86cd799439011")
// @Success      200 {object} domain.Response{data=string} "Successfully deleted bar"
// @Failure      400 {object} domain.Response "Bad request - Invalid ID format"
// @Failure      401 {object} domain.Response "Unauthorized - Missing or invalid token"
// @Failure      404 {object} domain.Response "Bar not found"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /bar/{id} [delete]
func (bh *BarHandler) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	if err := bh.barSvc.DeleteById(c.Request.Context(), id); err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "bar not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Bar deleted successfully", ""))
}
