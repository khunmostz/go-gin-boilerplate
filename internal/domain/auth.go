package domain

// LoginRequest represents the login request payload
// @Description User login request with email and password credentials
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@example.com" description:"User email address (must be valid email format)"`
	Password string `json:"password" binding:"required" example:"password123" description:"User password (minimum 6 characters)"`
}

// RegisterRequest represents the registration request payload
// @Description User registration request with email and password
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"newuser@example.com" description:"User email address (must be valid email format)"`
	Password string `json:"password" binding:"required" example:"securepassword123" description:"User password (minimum 6 characters)"`
}

// LoginResponse represents the login response with JWT tokens
// @Description Successful login response containing user information and JWT tokens
type LoginResponse struct {
	Email        string `json:"email" example:"admin@example.com" description:"Authenticated user's email address"`
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." description:"JWT access token for API authentication"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." description:"JWT refresh token for obtaining new access tokens"`
}
