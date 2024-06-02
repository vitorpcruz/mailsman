package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/mailsman/api/handler"
	"github.com/mailsman/internal/use_case/create_batch"
)

func main() {
	healthCheckerHandler := handler.InitHandler()

	batchEmailUseCase := create_batch.InstantiateUseCase()
	emailBatchHandler := handler.NewEmailBatchHandler(batchEmailUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", healthCheckerHandler.GetHealth)
	r.Post("/email/new", emailBatchHandler.CreateBatch)

	http.ListenAndServe(":3000", r)
}
