package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type ResponseContent struct {
	Response
	Content interface{} `json:"content"`
}

func New(msg string, httpCode int, w http.ResponseWriter) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(Response{Message: msg})
}

func NewWithContent(msg string, httpCode int, content interface{}, w http.ResponseWriter) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(
		ResponseContent{
			Response: Response{Message: msg},
			Content:  content,
		},
	)
}
