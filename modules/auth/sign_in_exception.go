package auth

import (
	"net/http"

	"github.com/oh-jinsu/helloworld/modules/common"
)

const (
	USER_NOT_FOUND = 1101 + iota
	PASSWORD_NOT_MATCHED
	FAILED_TO_ISSUE_ACCESS_TOKEN
	FAILED_TO_ISSUE_REFRESH_TOKEN
)

func UserNotFoundException() common.Exception {
	return common.NewError(http.StatusNotFound, USER_NOT_FOUND, "이용자를 찾지 못했습니다.")
}

func PasswordNotMatchedException() common.Exception {
	return common.NewError(http.StatusUnauthorized, PASSWORD_NOT_MATCHED, "비밀번호가 틀렸습니다.")
}

func FailedToIssueAccessTokenException() common.Exception {
	return common.NewError(http.StatusInternalServerError, FAILED_TO_ISSUE_ACCESS_TOKEN, "액세스 토큰을 발급하지 못했습니다.")
}

func FailedToIssueRefreshTokenException() common.Exception {
	return common.NewError(http.StatusInternalServerError, FAILED_TO_ISSUE_REFRESH_TOKEN, "리프레시 토큰을 발급하지 못했습니다.")
}
