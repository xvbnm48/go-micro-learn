package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
)

// routes sets up the routes for the API
func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()
	// specify who is allowed
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// middleware
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Post("/", app.Broker)
	return mux
}
