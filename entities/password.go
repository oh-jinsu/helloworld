package entities

import "errors"

type Password struct {
	value string
}

func NewPassword(value string) (*Password, *Exception) {
	if hasSpaceCharacters(value) {
		return nil, NewException(errors.New("비밀번호에 공백을 입력할 수 없습니다"))
	}

	if byteLen(value) < 8 {
		return nil, NewException(errors.New("비밀번호가 너무 짧습니다"))
	}

	if byteLen(value) > 24 {
		return nil, NewException(errors.New("비밀번호가 너무 깁니다"))
	}

	return &Password{value}, NewException(nil)
}

func CopyPassword(value string) *Password {
	return &Password{value}
}

func (e *Password) ToString() string {
	return e.value
}

func (e *Password) Equals(other *Password) bool {
	return e.value == other.value
}
