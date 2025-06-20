package handler

import (
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FooHandler struct {
	fooSvc port.FooService
}

func NewFooHandler(fooSvc port.FooService) *FooHandler {
	return &FooHandler{fooSvc: fooSvc}
}

// Create creates a new foo item
// @Summary      Create a new foo
// @Description  Create a new foo item with provided information
// @Tags         foo
// @Accept       json
// @Produce      json
// @Param        foo body domain.Foo true "Foo creation data"
// @Success      201 {object} domain.Response{data=domain.Foo} "Successfully created foo"
// @Failure      400 {object} domain.Response "Bad request - Invalid input data"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /foo [post]
func (fh *FooHandler) Create(c *gin.Context) {
	foo := &domain.Foo{}
	if err := c.ShouldBindJSON(foo); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse(err.Error()))
		return
	}

	createdFoo, err := fh.fooSvc.Create(c.Request.Context(), foo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponseWithMessage("Foo created successfully", createdFoo))
}

// GetAll retrieves all foo items
// @Summary      Get all foos
// @Description  Retrieve all foo items from the database
// @Tags         foo
// @Produce      json
// @Success      200 {object} domain.Response{data=[]domain.Foo} "Successfully retrieved all foos"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /foo [get]
func (fh *FooHandler) GetAll(c *gin.Context) {
	foos, err := fh.fooSvc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Foos retrieved successfully", foos))
}

// GetByID retrieves a foo by ID
// @Summary      Get foo by ID
// @Description  Retrieve a specific foo item by its unique identifier
// @Tags         foo
// @Produce      json
// @Param        id path string true "Foo ID" example("507f1f77bcf86cd799439011")
// @Success      200 {object} domain.Response{data=domain.Foo} "Successfully retrieved foo"
// @Failure      400 {object} domain.Response "Bad request - Invalid ID format"
// @Failure      404 {object} domain.Response "Foo not found"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /foo/{id} [get]
func (fh *FooHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	foo, err := fh.fooSvc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Foo retrieved successfully", foo))
}

// UpdateByID updates a foo by ID
// @Summary      Update foo by ID
// @Description  Update a specific foo item by its unique identifier
// @Tags         foo
// @Accept       json
// @Produce      json
// @Param        id path string true "Foo ID" example("507f1f77bcf86cd799439011")
// @Param        update body map[string]interface{} true "Update data" example({"name": "Updated Foo Name"})
// @Success      200 {object} domain.Response{data=domain.Foo} "Successfully updated foo"
// @Failure      400 {object} domain.Response "Bad request - Invalid input data or ID format"
// @Failure      404 {object} domain.Response "Foo not found"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /foo/{id} [put]
func (fh *FooHandler) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var update map[string]any
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse(err.Error()))
		return
	}

	updatedFoo, err := fh.fooSvc.UpdateById(c.Request.Context(), id, update)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "foo not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Foo updated successfully", updatedFoo))
}

// DeleteByID deletes a foo by ID
// @Summary      Delete foo by ID
// @Description  Delete a specific foo item by its unique identifier
// @Tags         foo
// @Produce      json
// @Param        id path string true "Foo ID" example("507f1f77bcf86cd799439011")
// @Success      200 {object} domain.Response{data=string} "Successfully deleted foo"
// @Failure      400 {object} domain.Response "Bad request - Invalid ID format"
// @Failure      404 {object} domain.Response "Foo not found"
// @Failure      500 {object} domain.Response "Internal server error"
// @Router       /foo/{id} [delete]
func (fh *FooHandler) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	if err := fh.fooSvc.DeleteById(c.Request.Context(), id); err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "foo not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Foo deleted successfully", ""))
}
