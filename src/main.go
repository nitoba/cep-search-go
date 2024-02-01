package main

import (
	"cep-search/src/infra/http"
)

func main() {
	cepSearchAPI := http.CepSearchAPI{}
	cepSearchAPI.Handle()
}
