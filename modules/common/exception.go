package common

import "net/http"

type Exception interface {
	Status() int
	Code() int
	Message() string
}

type exception struct {
	status  int
	code    int
	message string
}

func NewException(status int, code int, message string) Exception {
	return &exception{status, code, message}
}

func (e *exception) Status() int {
	return e.status
}

func (e *exception) Code() int {
	return e.code
}

func (e *exception) Message() string {
	return e.message
}

func BadRequestException() Exception {
	return &exception{http.StatusBadRequest, 1, "유효하지 않은 요청입니다."}
}
