package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)
	//resultado, services := utils.ListServiceProjeto(token, url, "teste-01")
	resultado, services := utils.ListService(token, url)
	if resultado == 0 {
		for i := 0; i < len(services.Items); i++ {
			fmt.Printf("Dc: %+v\r\n", services.Items[i].Metadata.Name)
		}
	} else {
		fmt.Println("[main] Services não encontrados em ", url)
	}
}
