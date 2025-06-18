package handler

import (
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"go-gin-boilerplate/internal/utils"
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
// @Summary      User login
// @Description  Authenticate user with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginRequest body domain.LoginRequest true "Login credentials"
// @Success      200 {object} utils.Response "Login successful"
// @Failure      400 {object} utils.Response "Bad request"
// @Failure      401 {object} utils.Response "Unauthorized"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq domain.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	loginResp, err := h.authSvc.Login(loginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponseWithMessage("Login successful", loginResp))
}
