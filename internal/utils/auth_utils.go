package utils

import (
	"errors"

	"github.com/Eugene600/Go-Project/internal/config"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func VerifyJwt(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(config.Cfg.Auth.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token after token parsing")
	}

	return token, nil
}

func RetrieveJwtClaim(claims jwt.Claims, value string) (any, bool) {
	extClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}
	val, ok := extClaims[value]

	return val, ok
}

func GetUserIdFromToken(token *jwt.Token) (uuid.UUID, error) {
	subClaim, ok := RetrieveJwtClaim(token.Claims, "sub")
	if !ok || subClaim == nil {
		return uuid.Nil, errors.New("sub claim not found or is not Ok")
	}

	sub, ok := subClaim.(string)
	if !ok {
		return uuid.Nil, errors.New("sub claim is not a string")
	}

	userId, err := uuid.FromString(sub)

	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}
