package resp

import "net/http"

func NotFoundError(message string) *Response {
	return &Response{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func UnexpectedError(message string) *Response {
	return &Response{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
