package auth

import (
	"net/http"

	"github.com/oh-jinsu/helloworld/modules/common"
)

const (
	PASSWORD_NOT_MATCHED = 10201 + iota
	FAILED_TO_ISSUE_REFRESH_TOKEN
)

func PasswordNotMatchedException() common.Exception {
	return common.NewException(http.StatusUnauthorized, PASSWORD_NOT_MATCHED, "비밀번호가 틀렸습니다.")
}

func FailedToIssueRefreshTokenException() common.Exception {
	return common.NewException(http.StatusInternalServerError, FAILED_TO_ISSUE_REFRESH_TOKEN, "리프레시 토큰을 발급하지 못했습니다.")
}
