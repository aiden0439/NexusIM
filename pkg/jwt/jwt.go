package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/aiden0439/NexusIM/pkg/config"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint64, cfg config.JWTConfig) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.ExpireHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.Secret))
}

func ParseToken(tokenString string, cfg config.JWTConfig) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{}, 
		func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
