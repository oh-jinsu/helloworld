package auth

import (
	"net/http"

	"github.com/oh-jinsu/helloworld/modules/common"
)

const (
	USER_NOT_FOUND = 10001 + iota
	FAILED_TO_ISSUE_ACCESS_TOKEN
)

func UserNotFoundException() common.Exception {
	return common.NewException(http.StatusNotFound, USER_NOT_FOUND, "이용자를 찾지 못했습니다.")
}

func FailedToIssueAccessTokenException() common.Exception {
	return common.NewException(http.StatusInternalServerError, FAILED_TO_ISSUE_ACCESS_TOKEN, "액세스 토큰을 발급하지 못했습니다.")
}
