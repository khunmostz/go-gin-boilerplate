package service

import (
	"errors"
	"go-gin-boilerplate/internal/domain"
	"go-gin-boilerplate/internal/port"
	"go-gin-boilerplate/internal/utils"
)

type AuthService struct {
	jwt *utils.JWT
}

func NewAuthService(jwt *utils.JWT) port.AuthService {
	return &AuthService{
		jwt: jwt,
	}
}

// Login is a mock implementation that validates against hardcoded users
func (s *AuthService) Login(req domain.LoginRequest) (domain.LoginResponse, error) {
	// Mock users for testing
	mockUsers := map[string]string{
		"admin@example.com": "password123",
		"user@example.com":  "password456",
		"test@example.com":  "test123",
	}

	// Check if user exists and password is correct
	if password, exists := mockUsers[req.Email]; !exists || password != req.Password {
		return domain.LoginResponse{}, errors.New("invalid email or password")
	}

	// Generate real JWT tokens
	accessToken, err := s.jwt.GenerateAccessToken(req.Email)
	if err != nil {
		return domain.LoginResponse{}, errors.New("failed to generate access token")
	}

	refreshToken, err := s.jwt.GenerateRefreshToken(req.Email)
	if err != nil {
		return domain.LoginResponse{}, errors.New("failed to generate refresh token")
	}

	return domain.LoginResponse{
		Email:        req.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
