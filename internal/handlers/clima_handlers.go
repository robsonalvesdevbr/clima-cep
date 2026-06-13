package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/dto"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/entity"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/infra/database"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/pkg"
)

type ClimaHandler struct {
	cepRepository   entity.CepRepositoryInterface
	climaRepository entity.ClimaRepositoryInterface
}

func NewClimaHandler(cepRepository entity.CepRepositoryInterface, climaRepository entity.ClimaRepositoryInterface) *ClimaHandler {
	return &ClimaHandler{
		cepRepository:   cepRepository,
		climaRepository: climaRepository,
	}
}

func (ch *ClimaHandler) ClimaHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	result, err := ch.cepRepository.GetCEP(cep)
	if err != nil {
		if err == database.CEPInvalidoError {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if err == database.CEPNaoEncontradoError {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clima, err := ch.climaRepository.GetClima(cep, result.Localidade)
	if err != nil {
		if err == database.CEPNaoEncontradoError {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clima.WeatherAPI)
}

func (ch *ClimaHandler) TempHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	result, err := ch.cepRepository.GetCEP(cep)
	if err != nil {
		if err == database.CEPInvalidoError {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if err == database.CEPNaoEncontradoError {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clima, err := ch.climaRepository.GetClima(cep, result.Localidade)
	if err != nil {
		if err == database.CEPNaoEncontradoError {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temps := dto.ClimaCEPDTO{
		TempC: clima.WeatherAPI.Current.Temp_c,
		TempF: pkg.CelsiusToFahrenheit(clima.WeatherAPI.Current.Temp_c),
		TempK: pkg.CelsiusToKelvin(clima.WeatherAPI.Current.Temp_c),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temps)
}
