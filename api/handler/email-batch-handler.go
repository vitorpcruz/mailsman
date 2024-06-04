package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mailsman/internal/use_case/create_batch"
	"github.com/mailsman/pkg/responses"
)

var response *responses.Response

type EmailBatchHandler struct {
	createBatchUseCase create_batch.UseCaseInterface
}

func InstantiateBatchHandler(useCase create_batch.UseCaseInterface) *EmailBatchHandler {
	return &EmailBatchHandler{createBatchUseCase: useCase}
}

func (handler *EmailBatchHandler) OnSucess(message string) {
	response = &responses.Response{Message: message}
}

func (handler *EmailBatchHandler) OnFail(err error) {
	response = &responses.Response{Message: err.Error()}
}

func (handler *EmailBatchHandler) CreateBatch(w http.ResponseWriter, r *http.Request) {
	var input create_batch.Input
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		handler.OnFail(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	handler.createBatchUseCase.Execute(input, handler)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
