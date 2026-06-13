package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/handlers"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/infra/database"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	viaCep := database.NewViaCepRepository()
	handlerCep := handlers.NewCepHandler(viaCep)

	weatherApi := database.NewWeatherAPIRepository()
	climaHandler := handlers.NewClimaHandler(viaCep, weatherApi)

	router.Route("/", func(r chi.Router) {
		r.Get("/", handlers.HelloWorldHandler)
	})

	router.Route("/hello", func(r chi.Router) {
		r.Get("/", handlers.HelloWorldHandler)
	})

	router.Route("/clima", func(r chi.Router) {
		r.Get("/", climaHandler.ClimaHandler)
		r.Get("/temp", climaHandler.TempHandler)
	})

	router.Route("/cep", func(r chi.Router) {
		r.Get("/", handlerCep.CepHandler)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}
