package entities

import "regexp"

type Password struct {
	value string
}

func NewPassword(value string) *Password {
	return &Password{value}
}

func (e *Password) ToString() string {
	return e.value
}

func (e *Password) Equals(other *Password) bool {
	return e.value == other.value
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
