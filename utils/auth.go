package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, secretKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()
	claims["sub"] = email
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(jwtToken string, secretKey []byte) (*jwt.MapClaims, error) {
	jwtTokenSanitised := strings.Replace(jwtToken, "Bearer ", "", 1)

	token, err := jwt.Parse(jwtTokenSanitised, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("Token not valid")
}
