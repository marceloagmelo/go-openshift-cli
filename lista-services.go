package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL")
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
