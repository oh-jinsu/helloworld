package auth

import (
	"net/http"

	"github.com/oh-jinsu/helloworld/modules/common"
)

const (
	CONFLICT_USERNAME = 1001 + iota
	KOREAN_CONSONANTS
	SPECIAL_CHARACTER
	SPACE_CHARACTER_FOR_USERNAME
	TOO_SHORT_USERNAME
	TOO_LONG_USERNAME
	TOO_SHORT_PASSWORD
	TOO_LONG_PASSWORD
	SPACE_CHARACTER_FOR_PASSWORD
	FAILED_TO_FIND_USER
)

func ConflictUsernameException() common.Exception {
	return common.NewException(http.StatusConflict, CONFLICT_USERNAME, "이미 존재하는 이름입니다.")
}

func KoreanConsonantsException() common.Exception {
	return common.NewException(http.StatusBadRequest, KOREAN_CONSONANTS, "이름에 한글 자음만 입력할 수 없습니다.")
}

func SpecialCharacterException() common.Exception {
	return common.NewException(http.StatusBadRequest, SPECIAL_CHARACTER, "이름에 특수문자를 입력할 수 없습니다.")
}

func SpaceCharacterForUsernameException() common.Exception {
	return common.NewException(http.StatusBadRequest, SPACE_CHARACTER_FOR_USERNAME, "이름에 공백을 입력할 수 없습니다.")
}

func TooShortUsernameException() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_SHORT_USERNAME, "이름이 너무 짧습니다.")
}

func TooLongUsernameException() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_LONG_USERNAME, "이름이 너무 깁니다.")
}

func TooShortPasswordException() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_SHORT_PASSWORD, "비밀번호가 너무 짧습니다.")
}

func TooLongPasswordException() common.Exception {
	return common.NewException(http.StatusBadRequest, TOO_LONG_PASSWORD, "비밀번호가 너무 깁니다.")
}

func SpaceCharacterForPasswordException() common.Exception {
	return common.NewException(http.StatusBadRequest, SPACE_CHARACTER_FOR_PASSWORD, "비밀번호에 공백을 입력할 수 없습니다.")
}

func FailedToFindUserException() common.Exception {
	return common.NewException(http.StatusInternalServerError, FAILED_TO_FIND_USER, "저장한 유저를  찾지 못했습니다.")
}
