package entities

import "regexp"

type Password struct {
	Value string
}

func (e *Password) ToString() string {
	return e.Value
}

func (e *Password) Equals(other *Password) bool {
	return e.Value == other.Value
}

func (e *Password) HasSpaceCharacters() bool {
	result, _ := regexp.MatchString("\\s", e.Value)

	return result
}

func (e *Password) IsTooShort() bool {
	return len([]byte(e.Value)) < 8
}

func (e *Password) IsTooLong() bool {
	return len([]byte(e.Value)) > 24
}
