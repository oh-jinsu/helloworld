package entities

import "regexp"

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
