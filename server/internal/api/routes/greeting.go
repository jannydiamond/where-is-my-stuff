package routes

import (
	"wims/internal/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetHelloRoutes(r chi.Router) {
	r.Get("/hello", handlers.GetHelloHandler)
	r.Get("/hello/{name}", handlers.GetHelloWithParamHandler)
}
