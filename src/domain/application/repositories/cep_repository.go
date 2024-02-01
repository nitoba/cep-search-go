package repositories

import . "cep-search/src/domain/enterprise/entities"

type CepRepository interface {
	GetAddress(cep string) (*Address, error)
}
