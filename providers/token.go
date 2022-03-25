package providers

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oh-jinsu/helloworld/entities"
)

func issueJwtToken(secret *[]byte, claims *entities.Claims) (string, error) {
	jwtClaims := jwt.MapClaims{}

	jwtClaims["authorized"] = true

	jwtClaims["iss"] = claims.Issuer()

	jwtClaims["user_id"] = strconv.Itoa(int(claims.UserId()))

	jwtClaims["exp"] = claims.Expiration().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	result, err := token.SignedString(*secret)

	if err != nil {
		return "", err
	}

	return result, nil
}

func extractClaims(token *jwt.Token) (*entities.Claims, error) {
	if !token.Valid {
		return &entities.Claims{}, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return &entities.Claims{}, nil
	}

	userId, err := strconv.ParseUint(claims["user_id"].(string), 10, 32)

	if err != nil {
		return &entities.Claims{}, nil
	}

	issuer := claims["iss"].(string)

	expiration, err := strconv.ParseInt(claims["exp"].(string), 10, 32)

	if err != nil {
		return &entities.Claims{}, nil
	}

	return entities.NewClaims(uint(userId), issuer, time.Unix(expiration, 0)), nil
}

func VerifyAccessToken(token string) (*entities.Claims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("")
		}

		return []byte(os.Getenv("JWT_SECRET_ACCESS_TOKEN")), nil
	})

	if err != nil {
		return &entities.Claims{}, nil
	}

	return extractClaims(t)
}

func IssueAccessToken(claims *entities.Claims) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET_ACCESS_TOKEN"))

	return issueJwtToken(&secret, claims)
}

func VerifyRefreshToken(token string) (*entities.Claims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("")
		}

		return []byte(os.Getenv("JWT_SECRET_REFRESH_TOKEN")), nil
	})

	if err != nil {
		return &entities.Claims{}, nil
	}

	return extractClaims(t)
}

func IssueRefreshToken(claims *entities.Claims) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET_REFRESH_TOKEN"))

	return issueJwtToken(&secret, claims)
}
