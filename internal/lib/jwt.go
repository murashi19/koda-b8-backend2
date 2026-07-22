package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTPayload struct{
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

const (
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 3 * 24 * time.Hour
)

func GenerateAccessToken(userID int64) (string, error) {
	claims := JWTPayload{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken(userID int64) (string, error) {
	claims := JWTPayload{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func HashRefreshToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func VerifyAccessToken(token string) (bool, *int64) {
	payload, err := jwt.ParseWithClaims(token, &JWTPayload{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false, nil
	}
	result, ok := payload.Claims.(*JWTPayload)
	if !ok || !payload.Valid {
		return false, nil
	}
	return true, &result.UserID
}

func VerifyRefreshToken(token string) (bool, *int64) {
	return VerifyAccessToken(token)
}