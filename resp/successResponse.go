package resp

import "net/http"

func ContinueSuccess(message string) *Response {
	return &Response{
		Message: message,
		Code:    http.StatusContinue,
	}
}

func OKSuccess(message string, data interface{}) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
