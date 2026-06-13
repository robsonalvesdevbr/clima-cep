package entity

type CepRepositoryInterface interface {
	GetCEP(cep string) (ClimaCEP, error)
}

type ClimaRepositoryInterface interface {
	GetClima(cep string, city string) (ClimaCEP, error)
}
