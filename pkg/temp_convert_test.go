package pkg

import (
	"math"
	"testing"
)

const tolerance = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name    string
		celsius float64
		want    float64
	}{
		{"exemplo do contrato", 28.5, 83.3},
		{"zero", 0, 32},
		{"ponto de ebulicao", 100, 212},
		{"negativo", -40, -40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tt.celsius)
			if !almostEqual(got, tt.want) {
				t.Errorf("CelsiusToFahrenheit(%v) = %v; want %v", tt.celsius, got, tt.want)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name    string
		celsius float64
		want    float64
	}{
		{"exemplo do contrato", 28.5, 301.65},
		{"zero", 0, 273.15},
		{"zero absoluto", -273.15, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToKelvin(tt.celsius)
			if !almostEqual(got, tt.want) {
				t.Errorf("CelsiusToKelvin(%v) = %v; want %v", tt.celsius, got, tt.want)
			}
		})
	}
}
