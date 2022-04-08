package resp

import "net/http"

type AppSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ContinueSuccess(message string, data interface{}) *AppSuccess {
	return &AppSuccess{
		Code:    http.StatusContinue,
		Message: message,
		Data:    data,
	}
}

func OKSuccess(message string, data interface{}) *AppSuccess {
	return &AppSuccess{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
