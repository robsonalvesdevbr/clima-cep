package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/dto"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/entity"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/infra/database"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

// mockCepRepository implementa entity.CepRepositoryInterface.
type mockCepRepository struct {
	result entity.ClimaCEP
	err    error
}

func (m *mockCepRepository) GetCEP(cep string) (entity.ClimaCEP, error) {
	return m.result, m.err
}

// mockClimaRepository implementa entity.ClimaRepositoryInterface.
type mockClimaRepository struct {
	result entity.ClimaCEP
	err    error
}

func (m *mockClimaRepository) GetClima(cep string, city string) (entity.ClimaCEP, error) {
	return m.result, m.err
}

func TestTempHandler_Sucesso(t *testing.T) {
	cepRepo := &mockCepRepository{
		result: entity.ClimaCEP{ViaCep: entity.ViaCep{Localidade: "Curitiba"}},
	}
	climaRepo := &mockClimaRepository{
		result: entity.ClimaCEP{
			WeatherAPI: entity.WeatherAPI{Current: entity.Current{Temp_c: 28.5}},
		},
	}

	handler := NewClimaHandler(cepRepo, climaRepo)

	req := httptest.NewRequest(http.MethodGet, "/clima/temp?cep=80050250", nil)
	rec := httptest.NewRecorder()
	handler.TempHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d; want %d", rec.Code, http.StatusOK)
	}

	var body dto.ClimaCEPDTO
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("falha ao decodificar resposta: %v", err)
	}

	if !almostEqual(body.TempC, 28.5) {
		t.Errorf("temp_C = %v; want 28.5", body.TempC)
	}
	if !almostEqual(body.TempF, 83.3) {
		t.Errorf("temp_F = %v; want 83.3", body.TempF)
	}
	if !almostEqual(body.TempK, 301.65) {
		t.Errorf("temp_K = %v; want 301.65", body.TempK)
	}
}

func TestTempHandler_CEPInvalido(t *testing.T) {
	cepRepo := &mockCepRepository{err: database.CEPInvalidoError}
	climaRepo := &mockClimaRepository{}

	handler := NewClimaHandler(cepRepo, climaRepo)

	req := httptest.NewRequest(http.MethodGet, "/clima/temp?cep=9995025", nil)
	rec := httptest.NewRecorder()
	handler.TempHandler(rec, req)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d; want %d", rec.Code, http.StatusUnprocessableEntity)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "invalid zipcode" {
		t.Errorf("body = %q; want %q", got, "invalid zipcode")
	}
}

func TestTempHandler_CEPNaoEncontrado(t *testing.T) {
	cepRepo := &mockCepRepository{err: database.CEPNaoEncontradoError}
	climaRepo := &mockClimaRepository{}

	handler := NewClimaHandler(cepRepo, climaRepo)

	req := httptest.NewRequest(http.MethodGet, "/clima/temp?cep=99950250", nil)
	rec := httptest.NewRecorder()
	handler.TempHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d; want %d", rec.Code, http.StatusNotFound)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "can not find zipcode" {
		t.Errorf("body = %q; want %q", got, "can not find zipcode")
	}
}
