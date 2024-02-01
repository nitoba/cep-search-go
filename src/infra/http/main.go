package http

import (
	"cep-search/src/infra/http/controllers"
	"fmt"
	"net/http"
)

type CepSearchAPI struct{}

func (a *CepSearchAPI) Handle() {
	fetchCepController := controllers.CreateFetchCepController()
	mux := http.NewServeMux()

	mux.HandleFunc("/cep", fetchCepController.Handle)

	fmt.Println("Server listening on PORT: 3333")
	http.ListenAndServe(":3333", mux)
}
