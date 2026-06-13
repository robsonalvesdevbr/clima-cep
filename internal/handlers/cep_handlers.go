package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/entity"
	"github.com/robsonalvesdevbr/clima-cep-fullcycle/internal/infra/database"
)

type CepHandler struct {
	cep entity.CepRepositoryInterface
}

func NewCepHandler(cep entity.CepRepositoryInterface) *CepHandler {
	return &CepHandler{
		cep: cep,
	}
}

func (ch *CepHandler) CepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	result, err := ch.cep.GetCEP(cep)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
