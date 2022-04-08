package resp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func writeResponse(w http.ResponseWriter, r Response) {
	w.Header().Add("Content-Type", "application/json")
	code := 200
	if r.Code <= 400 {
		code = r.Code
	}
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		panic(err)
	}
}
