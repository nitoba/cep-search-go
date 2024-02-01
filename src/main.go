package main

import (
	"cep-search/src/infra/http"
)

func main() {
	cepSearchCLI := http.CepSearchAPI{}
	cepSearchCLI.Handle()
}
