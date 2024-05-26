package responses

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

type responseContent struct {
	response
	Content interface{} `json:"content"`
}

func Response(msg string, httpCode int, w http.ResponseWriter) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response{Message: msg})
}

func ResponseWithContent(msg string, httpCode int, content interface{}, w http.ResponseWriter) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(
		responseContent{
			response: response{Message: msg},
			Content:  content,
		},
	)
}