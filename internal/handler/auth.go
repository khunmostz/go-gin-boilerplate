package handler

import (
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSvc port.AuthService
}

func NewAuthHandler(authSvc port.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

// Login authenticates a user and returns access tokens
// @Summary      User authentication
// @Description  Authenticate user with email and password, returns JWT access and refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginRequest body domain.LoginRequest true "User login credentials" example({"email": "admin@example.com", "password": "password123"})
// @Success      200 {object} domain.Response{data=domain.LoginResponse} "Login successful - Returns access and refresh tokens"
// @Failure      400 {object} domain.Response "Bad request - Invalid JSON format or missing required fields"
// @Failure      401 {object} domain.Response "Unauthorized - Invalid email or password"
// @Failure      500 {object} domain.Response "Internal server error - Failed to generate tokens"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq domain.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse(err.Error()))
		return
	}

	loginResp, err := h.authSvc.Login(loginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponseWithMessage("Login successful", loginResp))
}
