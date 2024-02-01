package repositories

import (
	. "cep-search/src/domain/enterprise/entities"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type viaCepRepository struct{}

func (r viaCepRepository) GetAddress(cep string) (*Address, error) {
	URL_TO_SEARCH_ADDRESS := "https://viacep.com.br/ws/" + strings.TrimSpace(cep) + "/json/"
	res, err := http.Get(URL_TO_SEARCH_ADDRESS)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("The cep: %s does not exist, please try again with another!", strings.TrimSpace(cep)))
	}

	var address Address
	err = json.NewDecoder(res.Body).Decode(&address)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &address, nil
}

func CreateViaCepRepository() *viaCepRepository {
	return &viaCepRepository{}
}
