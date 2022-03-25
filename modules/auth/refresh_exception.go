package auth

import (
	"net/http"

	"github.com/oh-jinsu/helloworld/modules/common"
)

const (
	INVALID_REFRESH_TOKEN = 10301 + iota
	EXPIRED_REFRESH_TOKEN
)

func InvalidRefreshTokenException() common.Exception {
	return common.NewException(http.StatusUnauthorized, INVALID_REFRESH_TOKEN, "유효하지 않은 인증 정보입니다.")
}

func DiscardedRefreshTokenException() common.Exception {
	return common.NewException(http.StatusUnauthorized, EXPIRED_REFRESH_TOKEN, "이미 폐기된 인증 정보입니다.")
}
