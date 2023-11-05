package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"wims/internal/api/handlers"
	"wims/internal/api/routes"
	"wims/internal/config"
	"wims/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load the configuration
	cfg := config.LoadConfig()

	fmt.Printf("Server will listen on port %d\n", cfg.Server.Port)

	// Set up the database connection
	db, dbErr := database.Connect()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	// Api routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Register handlers and routes
	r.Get("/", handlers.GetRootHandler)

	routes.SetHelloRoutes(r)

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	err := http.ListenAndServe(addr, r)

	// Error handling when server is closes
	// or any other error
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
