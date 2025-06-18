package port

import "go-gin-boilerplate/internal/domain"

type AuthService interface {
	Login(req domain.LoginRequest) (domain.LoginResponse, error)
	// Register(req domain.RegisterRequest) (domain.RegisterResponse, error)
	// Logout(token string) error
	// RefreshToken(token string) (domain.LoginResponse, error)
}
