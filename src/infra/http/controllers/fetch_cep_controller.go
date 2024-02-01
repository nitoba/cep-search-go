package controllers

import (
	. "cep-search/src/domain/application/use_cases"
	. "cep-search/src/infra/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type cepSearchController struct {
	fetchAddressUseCase FetchAddressUseCase
}

func CreateFetchCepController() cepSearchController {
	cepRepository := CreateViaCepRepository()
	fetchAddressUseCase := CreateFetchAddressUseCase(cepRepository)
	return cepSearchController{
		fetchAddressUseCase: fetchAddressUseCase,
	}
}

func (c *cepSearchController) Handle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	cep := req.URL.Query().Get("cep")
	if cep == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"message": "cep is required"}`))
		return
	}

	address, err := c.fetchAddressUseCase.Execute(cep)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(fmt.Sprintf(`{"message": "%v"}`, err.Error())))
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(address)
}
