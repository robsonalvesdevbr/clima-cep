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

func (r *ViaCepRepository) GetCEP(cep string) (any, error) {
	if !r.viaCep.ValidateCEP(cep) {
		return entity.ViaCep{}, CEPInvalidoError
	}

	u := url.URL{
		Scheme: "https",
		Host:   "viacep.com.br",
		Path:   fmt.Sprintf("/ws/%s/json/", cep),
	}

	req, err := http.Get(u.String())
	if err != nil {
		return entity.ViaCep{}, CEPNaoEncontradoError
	}
	defer req.Body.Close()

	var response struct {
		entity.ViaCep
		Erro json.RawMessage `json:"erro"`
	}

	if err := json.NewDecoder(req.Body).Decode(&response); err != nil {
		return entity.ViaCep{}, err
	}

	if isTrueJSONValue(response.Erro) {
		return entity.ViaCep{}, CEPNaoEncontradoError
	}

	return response.ViaCep, nil
}

func isTrueJSONValue(value json.RawMessage) bool {
	if len(value) == 0 {
		return false
	}

	return string(value) == "true" || string(value) == `"true"`
}
