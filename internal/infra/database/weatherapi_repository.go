package database

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/entity"
)

// defaultWeatherAPIKey é usada como fallback quando WEATHER_API_KEY não está definida.
const defaultWeatherAPIKey = "231c557276604a7e9b4181139261306"

func weatherAPIKey() string {
	if key := os.Getenv("WEATHER_API_KEY"); key != "" {
		return key
	}
	return defaultWeatherAPIKey
}

type WeatherAPIRepository struct {
	weatherApi entity.WeatherAPI
}

func NewWeatherAPIRepository() *WeatherAPIRepository {
	return &WeatherAPIRepository{
		weatherApi: *entity.NewWeatherAPI(),
	}
}

func (r *WeatherAPIRepository) GetClima(cep string, city string) (entity.ClimaCEP, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.weatherapi.com",
		Path:   "/v1/current.json",
		RawQuery: url.Values{
			"key": []string{weatherAPIKey()},
			"q":   []string{fmt.Sprintf("%s:%s", cep, city)},
			"aqi": []string{"no"},
		}.Encode(),
	}

	req, err := http.Get(u.String())
	if err != nil {
		return entity.ClimaCEP{}, CEPNaoEncontradoError
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return entity.ClimaCEP{}, CEPNaoEncontradoError
	}

	var response struct {
		entity.WeatherAPI
	}
	if err := json.NewDecoder(req.Body).Decode(&response); err != nil {
		return entity.ClimaCEP{}, err
	}

	return entity.ClimaCEP{
		WeatherAPI: response.WeatherAPI,
	}, nil
}
