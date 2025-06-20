package utils

import (
	"go-gin-boilerplate/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret          string
	AccessDuration  string
	RefreshDuration string
}

func NewJWT(config *config.JWTConfig) *JWT {
	return &JWT{
		Secret:          config.Secret,
		AccessDuration:  config.AccessDuration.String(),
		RefreshDuration: config.RefreshDuration.String(),
	}
}

func (j *JWT) GenerateAccessToken(userID string) (string, error) {
	accessDuration, err := time.ParseDuration(j.AccessDuration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(accessDuration).Unix(),
	})

	accessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (j *JWT) GenerateRefreshToken(userID string) (string, error) {
	refreshDuration, err := time.ParseDuration(j.RefreshDuration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(refreshDuration).Unix(),
	})

	refreshToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (j *JWT) ValidateToken(token string) (string, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return "", err
	}

	return claims["sub"].(string), nil
}
