package resp

import "net/http"

func ContinueSuccess(message string, data interface{}) *AppResponse {
	return &AppResponse{
		Code:    http.StatusContinue,
		Message: message,
		Data:    data,
	}
}

func OKSuccess(message string, data interface{}) *AppResponse {
	return &AppResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
