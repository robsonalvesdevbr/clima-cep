package entity

import "testing"

func TestValidateCEP(t *testing.T) {
	tests := []struct {
		name string
		cep  string
		want bool
	}{
		{"valido com 8 digitos", "01001000", true},
		{"valido outro", "80050250", true},
		{"menos de 8 digitos", "9995025", false},
		{"mais de 8 digitos", "999502501", false},
		{"vazio", "", false},
		{"com letras", "0100100a", false},
		{"com hifen", "0100-000", false},
		{"com espaco", "0100 000", false},
	}

	viaCep := NewViaCep()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := viaCep.ValidateCEP(tt.cep); got != tt.want {
				t.Errorf("ValidateCEP(%q) = %v; want %v", tt.cep, got, tt.want)
			}
		})
	}
}
