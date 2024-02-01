package use_cases

import (
	. "cep-search/src/domain/application/repositories"
	. "cep-search/src/domain/enterprise/entities"
)

type FetchAddressUseCase struct {
	cepRepository CepRepository
}

func (f *FetchAddressUseCase) Execute(cep string) (*Address, error) {
	address, err := f.cepRepository.GetAddress(cep)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func CreateFetchAddressUseCase(cepRepository CepRepository) FetchAddressUseCase {
	return FetchAddressUseCase{
		cepRepository: cepRepository,
	}
}
