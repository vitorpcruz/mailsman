package handler

import (
	"net/http"

	"github.com/mailsman/pkg/responses"
)

type HealthcheckHandler struct {
}

func InitHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	responses.New("It's running.", http.StatusOK, w)
}
