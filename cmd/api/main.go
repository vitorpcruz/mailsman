package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/mailsman/api/handler"
)

func main() {
	healthCheckerHandler := handler.InitHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", healthCheckerHandler.GetHealth)

	http.ListenAndServe(":3000", r)
}
