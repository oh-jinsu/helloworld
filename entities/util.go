package entities

import "regexp"

func hasKoreanConsonants(value string) bool {
	result, _ := regexp.MatchString("[ㄱ-ㅎ]", value)

	return result
}

func hasSpecialCharacters(value string) bool {
	result, _ := regexp.MatchString("[\\{\\}\\[\\]\\/?.,;:|\\)*~`!^\\-_+<>@\\#$%&\\\\\\=\\(\\'\"]", value)

	return result
}

func hasSpaceCharacters(value string) bool {
	result, _ := regexp.MatchString("\\s", value)

	return result
}

func byteLen(value string) int {
	return len([]byte(value))
}
