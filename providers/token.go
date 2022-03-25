package providers

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func issueJwtToken(secret *[]byte, claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString(*secret)

	if err != nil {
		return "", err
	}

	return result, nil
}

func extractClaims(token *jwt.Token) (*map[string]interface{}, error) {
	if !token.Valid {
		return &map[string]interface{}{}, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return &map[string]interface{}{}, nil
	}

	result := &map[string]interface{}{}

	for k, v := range claims {
		(*result)[k] = v
	}

	return result, nil
}

func VerifyAccessToken(token string) (*map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("")
		}

		return []byte(os.Getenv("JWT_SECRET_ACCESS_TOKEN")), nil
	})

	if err != nil {
		return &map[string]interface{}{}, nil
	}

	return extractClaims(t)
}

func IssueAccessToken(userId uint, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true

	claims["iss"] = os.Getenv("JWT_ISS")

	claims["user_id"] = strconv.Itoa(int(userId))

	claims["exp"] = time.Now().Add(exp).Unix()

	secret := []byte(os.Getenv("JWT_SECRET_ACCESS_TOKEN"))

	return issueJwtToken(&secret, &claims)
}

func VerifyRefreshToken(token string) (*map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("")
		}

		return []byte(os.Getenv("JWT_SECRET_REFRESH_TOKEN")), nil
	})

	if err != nil {
		return &map[string]interface{}{}, nil
	}

	return extractClaims(t)
}

func IssueRefreshToken(userId uint, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true

	claims["iss"] = os.Getenv("JWT_ISS")

	claims["user_id"] = strconv.Itoa(int(userId))

	claims["exp"] = time.Now().Add(exp).Unix()

	secret := []byte(os.Getenv("JWT_SECRET_REFRESH_TOKEN"))

	return issueJwtToken(&secret, &claims)
}
