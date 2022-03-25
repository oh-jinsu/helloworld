package entities

import (
	"time"
)

type User struct {
	id           uint
	username     *Username
	password     *Password
	refreshToken string
	createdAt    time.Time
}

func NewUserByUsernameAndPassword(username *Username, password *Password) *User {
	return &User{username: username, password: password}
}

func NewUser(id uint, username *Username, password *Password, refreshToken string, createdAt time.Time) *User {
	return &User{id, username, password, refreshToken, createdAt}
}

func (e *User) Id() uint {
	return e.id
}

func (e *User) Username() *Username {
	return e.username
}

func (e *User) Password() *Password {
	return e.password
}

func (e *User) RefreshToken() string {
	return e.refreshToken
}

func (e *User) CreatedAt() time.Time {
	return e.createdAt
}

func (e *User) UpdateRefreshToken(refreshToken string) {
	e.refreshToken = refreshToken
}
