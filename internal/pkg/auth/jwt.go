package auth

import (
	"codename-rl/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	ID string `json:"user_id"`
	jwt.RegisteredClaims
}

type JwtService struct {
	secret []byte
}

func NewJwtService(secret string) *JwtService {
	return &JwtService{secret: []byte(secret)}
}

func (r *JwtService) GenerateToken(user *entity.User) (string, error) {
	// Set the expiration time for the token (e.g., 24 hours).
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JwtClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "codename-rl",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(r.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *JwtService) ValidateToken(tokenString string) (*JwtClaims, error) {
	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return r.secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
