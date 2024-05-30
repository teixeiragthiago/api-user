package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtGeneratorService interface {
	GenerateToken(claimsData *Claims) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) JwtGeneratorService {
	return &jwtService{secretKey: secretKey}
}

type Claims struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func (j *jwtService) GenerateToken(claimsData *Claims) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:       claimsData.ID,
		Nickname: claimsData.Nickname,
		Email:    claimsData.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
