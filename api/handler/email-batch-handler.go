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

func InstantiateBatchHandler(useCase create_batch.UseCaseInterface) *EmailBatchHandler {
	return &EmailBatchHandler{createBatchUseCase: useCase}
}

func (handler *EmailBatchHandler) CreateBatch(w http.ResponseWriter, r *http.Request) {
	var input create_batch.Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		responses.New(err.Error(), http.StatusBadRequest, w)

		return
	}

	err = handler.createBatchUseCase.Execute(input)
	if err != nil {
		responses.New(err.Error(), http.StatusInternalServerError, w)

		return
	}

	responses.New("Email successfully sent", http.StatusOK, w)
}
