package entity

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCep() *ViaCep {
	return &ViaCep{}
}

func (viaCep *ViaCep) ValidateCEP(cep string) bool {
	if len(cep) != 8 {
		return false
	}

	for _, char := range cep {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}
