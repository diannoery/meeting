package message

import (
	"net/http"
)

type ResponeMessage struct {
	Message string
	Code    int
	Status  string
	Results interface{}
}

func Respone(msg string, code int, rslt interface{}) *ResponeMessage {

	return &ResponeMessage{
		Message: msg,
		Code:    code,
		Status:  http.StatusText(code),
		Results: rslt,
	}
}
