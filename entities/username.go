package entities

import "regexp"

type Username struct {
	Value string
}

func (e *Username) ToString() string {
	return e.Value
}

func (e *Username) HasKoreanConsonants() bool {
	result, _ := regexp.MatchString("[ㄱ-ㅎ]", e.Value)

	return result
}

func (e *Username) HasSpecialCharacters() bool {
	result, _ := regexp.MatchString("[\\{\\}\\[\\]\\/?.,;:|\\)*~`!^\\-_+<>@\\#$%&\\\\\\=\\(\\'\"]", e.Value)

	return result
}

func (e *Username) HasSpaceCharacters() bool {
	result, _ := regexp.MatchString("\\s", e.Value)

	return result
}

func (e *Username) IsTooShort() bool {
	return len([]byte(e.Value)) < 4
}

func (e *Username) IsTooLong() bool {
	return len([]byte(e.Value)) > 16
}
