package main

import (
	"go-react-auth/handler"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	router := chi.NewMux()
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	router.Get("/customer/{id}", handler.Make(handler.HandleGetCustomer))

	port := os.Getenv("LISTEN_ADDR")
	slog.Info("API server started", "port", os.Getenv("LISTEN_ADDR"))
	http.ListenAndServe(port, router)
}
