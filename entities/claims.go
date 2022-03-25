package entities

import (
	"os"
	"time"
)

type Claims struct {
	subject    string
	audience   string
	userId     uint
	issuer     string
	expiration time.Time
	issuedAt   time.Time
}

func NewClaims(subject string, userId uint, expiration time.Time) *Claims {
	return &Claims{subject, os.Getenv("ORIGIN"), userId, os.Getenv("JWT_ISSUER"), expiration, time.Now()}
}

func CopyClaims(subject string, audience string, userId uint, issuer string, expiration time.Time, issuedAt time.Time) *Claims {
	return &Claims{subject, audience, userId, issuer, expiration, issuedAt}
}

func (e *Claims) Subject() string {
	return e.subject
}

func (e *Claims) Audience() string {
	return e.audience
}

func (e *Claims) UserId() uint {
	return e.userId
}

func (e *Claims) Issuer() string {
	return e.issuer
}

func (e *Claims) Expiration() time.Time {
	return e.expiration
}

func (e *Claims) IssuedAt() time.Time {
	return e.issuedAt
}
