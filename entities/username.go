package entities

import "errors"

type Username struct {
	value string
}

func NewUsername(value string) (*Username, *Exception) {
	if hasKoreanConsonants(value) {
		return nil, NewException(errors.New("이름에 한글 자음만을 입력할 수 없습니다"))
	}

	if hasSpecialCharacters(value) {
		return nil, NewException(errors.New("이름에 특수문자를 입력할 수 없습니다"))
	}

	if hasSpaceCharacters(value) {
		return nil, NewException(errors.New("이름에 공백을 입력할 수 없습니다"))
	}

	if byteLen(value) < 4 {
		return nil, NewException(errors.New("이름이 너무 짧습니다"))
	}

	if byteLen(value) > 16 {
		return nil, NewException(errors.New("이름이 너무 깁니다"))
	}

	return &Username{value}, NewException(nil)
}

func CopyUsername(value string) *Username {
	return &Username{value}
}

func (e *Username) ToString() string {
	return e.value
}
