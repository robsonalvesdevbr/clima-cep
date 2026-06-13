package database

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/entity"
)

var (
	CEPInvalidoError      = fmt.Errorf("invalid zipcode")
	CEPNaoEncontradoError = fmt.Errorf("can not find zipcode")
)

type ViaCepRepository struct {
	viaCep entity.ViaCep
}

func NewViaCepRepository() *ViaCepRepository {
	return &ViaCepRepository{
		viaCep: *entity.NewViaCep(),
	}
}

func (r *ViaCepRepository) GetCEP(cep string) (entity.ClimaCEP, error) {
	if !r.viaCep.ValidateCEP(cep) {
		return entity.ClimaCEP{}, CEPInvalidoError
	}

	u := url.URL{
		Scheme: "https",
		Host:   "viacep.com.br",
		Path:   fmt.Sprintf("/ws/%s/json/", cep),
	}

	req, err := http.Get(u.String())
	if err != nil {
		return entity.ClimaCEP{}, CEPNaoEncontradoError
	}
	defer req.Body.Close()

	var response struct {
		entity.ViaCep
		Erro json.RawMessage `json:"erro"`
	}

	if err := json.NewDecoder(req.Body).Decode(&response); err != nil {
		return entity.ClimaCEP{}, err
	}

	if isTrueJSONValue(response.Erro) {
		return entity.ClimaCEP{}, CEPNaoEncontradoError
	}

	return entity.ClimaCEP{
		ViaCep: response.ViaCep,
	}, nil
}

func isTrueJSONValue(value json.RawMessage) bool {
	if len(value) == 0 {
		return false
	}

	return string(value) == "true" || string(value) == `"true"`
}
