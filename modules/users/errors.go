package users

import (
	"net/http"

	"github.com/oh-jinsu/helloworld/modules/common"
)

func ConflictUsernameException() common.Exception {
	return common.NewException(http.StatusConflict, 1001, "이미 존재하는 이름입니다.")
}
