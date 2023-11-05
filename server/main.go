package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"wims/internal/api/handlers"
	"wims/internal/config"
)

func main() {
	// Load the configuration
	cfg := config.LoadConfig()

	fmt.Printf("Server will listen on port %d\n", cfg.Server.Port)

	// Register handlers and routes
	http.HandleFunc("/", handlers.GetRootHandler)
	http.HandleFunc("/hello", handlers.GetHelloHandler)

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	err := http.ListenAndServe(addr, nil)

	// Error handling when server is closes
	// or any other error
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
