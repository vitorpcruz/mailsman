package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mailsman/internal/use_case/create_batch"
	"github.com/mailsman/pkg/responses"
)

type EmailBatchHandler struct {
	createBatchUseCase create_batch.UseCaseInterface
}

var msg string

func NewEmailBatchHandler(useCase create_batch.UseCaseInterface) *EmailBatchHandler {
	return &EmailBatchHandler{createBatchUseCase: useCase}
}

func (h *EmailBatchHandler) CreateBatch(w http.ResponseWriter, r *http.Request) {
	var input create_batch.Input
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		responses.New(msg, http.StatusBadRequest, w)
		return
	}

	h.createBatchUseCase.Execute(input, h)
	responses.New(msg, http.StatusOK, w)
}

func (h *EmailBatchHandler) OnSucess(message string) {
	msg = message
}

func (h *EmailBatchHandler) OnFail(message string) {
	msg = message
}
