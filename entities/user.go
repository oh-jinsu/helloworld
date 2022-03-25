package entities

import (
	"time"
)

type User struct {
	id        uint
	username  *Username
	password  *Password
	createdAt time.Time
}

func NewUser(username *Username, password *Password) *User {
	return &User{
		username: username,
		password: password,
	}
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

func (e *User) CreatedAt() time.Time {
	return e.createdAt
}
