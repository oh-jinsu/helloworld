package providers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oh-jinsu/helloworld/entities"
)

func IssueAccessToken(claims *entities.Claims) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET_ACCESS_TOKEN"))

	return issueJwtToken(&secret, claims)
}

func IssueRefreshToken(claims *entities.Claims) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET_REFRESH_TOKEN"))

	return issueJwtToken(&secret, claims)
}

func issueJwtToken(secret *[]byte, claims *entities.Claims) (string, error) {
	jwtClaims := jwt.MapClaims{}

	jwtClaims["sub"] = claims.Subject()

	jwtClaims["aud"] = claims.Audience()

	jwtClaims["iss"] = claims.Issuer()

	jwtClaims["user_id"] = claims.UserId()

	jwtClaims["exp"] = claims.Expiration().Unix()

	jwtClaims["iat"] = claims.IssuedAt().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	result, err := token.SignedString(*secret)

	if err != nil {
		return "", err
	}

	return result, nil
}

func VerifyAccessToken(token string) (*entities.Claims, error) {
	return verifyToken(token, []byte(os.Getenv("JWT_SECRET_ACCESS_TOKEN")))
}

func VerifyRefreshToken(token string) (*entities.Claims, error) {
	return verifyToken(token, []byte(os.Getenv("JWT_SECRET_REFRESH_TOKEN")))
}

func verifyToken(token string, secret []byte) (*entities.Claims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method is not matched")
		}

		return secret, nil
	})

	if err != nil || !t.Valid {
		return &entities.Claims{}, errors.New("given token is invalid")
	}

	claims, err := extractClaims(t)

	if err != nil {
		return &entities.Claims{}, err
	}

	ok := verifyClaims(claims)

	if !ok {
		return &entities.Claims{}, errors.New("given claims are dirty")
	}

	return claims, nil
}

func extractClaims(token *jwt.Token) (*entities.Claims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return &entities.Claims{}, errors.New("failed to cast claims")
	}

	subject, ok := claims["sub"].(string)

	if !ok {
		return &entities.Claims{}, errors.New("failed to cast claims")
	}

	audience, ok := claims["aud"].(string)

	if !ok {
		return &entities.Claims{}, errors.New("failed to cast claims")
	}

	userId, ok := claims["user_id"].(float64)

	if !ok {
		return &entities.Claims{}, errors.New("failed to cast claims")
	}

	issuer, ok := claims["iss"].(string)

	if !ok {
		return &entities.Claims{}, errors.New("failed to cast claims")
	}

	expiration, ok := claims["exp"].(float64)

	if !ok {
		return &entities.Claims{}, errors.New("failed to cast claims")
	}

	issuedAt := claims["iat"].(float64)

	if !ok {
		return &entities.Claims{}, errors.New("failed to cast claims")
	}

	return entities.CopyClaims(
		subject,
		audience,
		uint(userId),
		issuer,
		time.Unix(int64(expiration), 0),
		time.Unix(int64(issuedAt), 0),
	), nil
}

func verifyClaims(claims *entities.Claims) bool {
	if claims.Audience() != os.Getenv("JWT_AUDIENCE") {
		return false
	}

	if claims.Issuer() != os.Getenv("JWT_ISSUER") {
		return false
	}

	return true
}
