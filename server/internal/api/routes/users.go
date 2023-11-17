package routes

import (
	"database/sql"
	"net/http"
	"wims/internal/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetUsersRoutes(r chi.Router, db *sql.DB) {
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListUsersHandler(w, r, db)
	})
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUserHandler(w, r, db)
	})
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUserHandler(w, r, db)
	})
	r.Put("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUserHandler(w, r, db)
	})
	r.Delete("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUserHandler(w, r, db)
	})
}
