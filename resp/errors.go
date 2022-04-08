package resp

import "net/http"

func NotFoundError(message string) *AppResponse {
	return &AppResponse{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func UnexpectedError(message string) *AppResponse {
	return &AppResponse{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
