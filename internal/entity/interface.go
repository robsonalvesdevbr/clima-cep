package entity

type CepRepositoryInterface interface {
	GetCEP(cep string) (any, error)
}
