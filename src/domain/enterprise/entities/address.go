package entities

import (
	"fmt"
)

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func (a *Address) ShowAddress() string {
	return fmt.Sprintf("Your address contain the following information:\nCEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s", a.Cep, a.Logradouro, a.Bairro, a.Localidade, a.Uf)
}
