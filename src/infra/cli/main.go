package cli

import (
	"bufio"
	. "cep-search/src/domain/application/use_cases"
	. "cep-search/src/infra/repositories"
	"fmt"
	"os"
)

type CepSearchCLI struct{}

func (c *CepSearchCLI) Handle() {
	fmt.Println("Hello, this is a cep search")
	fmt.Println("Please enter your cep bellow:")

	typedCep, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Searching for this cep:%s", typedCep)

	cepRepository := CreateViaCepRepository()
	fetchAddressUseCase := CreateFetchAddressUseCase(cepRepository)

	address, err := fetchAddressUseCase.Execute(typedCep)

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("cep.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString(address.ShowAddress())
	fmt.Println(address.ShowAddress())
}
