package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func GetHelloWithParamHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
