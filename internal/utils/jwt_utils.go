package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/Eugene600/Go-Project/internal/config"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId.String(),
		"exp": time.Now().Add(time.Duration((config.Cfg.Auth.AccessTokenTTL) * int(time.Minute))).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.Cfg.Auth.JwtSecret))

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func GenerateRefreshToken(userId uuid.UUID) (string, time.Time, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", time.Time{}, err
	}

	refreshToken := base64.RawURLEncoding.EncodeToString(bytes)

	expiryDate := time.Now().AddDate(0, 0, config.Cfg.Auth.RefreshTokenTTL)

	return refreshToken, expiryDate, nil
}
