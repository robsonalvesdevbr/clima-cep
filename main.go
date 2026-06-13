package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/handlers"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/infra/database"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/hello", func(r chi.Router) {
		r.Get("/", handlers.HelloWorldHandler)
	})

	router.Route("/clima", func(r chi.Router) {
		r.Get("/", handlers.ClimaHandler)
	})

	viaCep := database.NewViaCepRepository()
	handlers := handlers.NewCepHandler(viaCep)

	router.Route("/cep", func(r chi.Router) {
		r.Get("/", handlers.CepHandler)
	})

	http.ListenAndServe(":8080", router)
}
