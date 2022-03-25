package entities

import "time"

type Claims struct {
	userId     uint
	issuer     string
	expiration time.Time
}

func NewClaims(userId uint, issuer string, expiration time.Time) *Claims {
	return &Claims{userId, issuer, expiration}
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
