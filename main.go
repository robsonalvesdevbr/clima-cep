package main

import (
	"fmt"
	"net/http"

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

	router.Route("/hello", func(r chi.Router) {
		r.Get("/", handlers.HelloWorldHandler)
	})

	router.Route("/clima", func(r chi.Router) {
		r.Get("/", climaHandler.ClimaHandler)
	})

	router.Route("/cep", func(r chi.Router) {
		r.Get("/", handlerCep.CepHandler)
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
