package entities

import (
	"regexp"

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

type Username struct {
	value string
}

func NewUsername(value string) *Username {
	return &Username{value}
}

func (e *Username) ToString() string {
	return e.value
}

func (e *Username) HasKoreanConsonants() bool {
	result, _ := regexp.MatchString("[ㄱ-ㅎ]", e.value)

	return result
}

func (e *Username) HasSpecialCharacters() bool {
	result, _ := regexp.MatchString("[\\{\\}\\[\\]\\/?.,;:|\\)*~`!^\\-_+<>@\\#$%&\\\\\\=\\(\\'\"]", e.value)

	return result
}

func (e *Username) HasSpaceCharacters() bool {
	result, _ := regexp.MatchString("\\s", e.value)

	return result
}

func (e *Username) IsTooShort() bool {
	return len([]byte(e.value)) < 4
}

func (e *Username) IsTooLong() bool {
	return len([]byte(e.value)) > 16
}

type Password struct {
	value string
}

func NewPassword(value string) *Password {
	return &Password{value}
}

func (e *Password) ToString() string {
	return e.value
}

func (e *Password) HasSpaceCharacters() bool {
	result, _ := regexp.MatchString("\\s", e.value)

	return result
}

func (e *Password) IsTooShort() bool {
	return len([]byte(e.value)) < 8
}

func (e *Password) IsTooLong() bool {
	return len([]byte(e.value)) > 24
}
