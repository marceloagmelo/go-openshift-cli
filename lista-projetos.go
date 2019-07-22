package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
        url := utils.URLGen("pre")
        token := utils.GetToken(url)
	resultado, projetos := utils.Namespaces(token, url)
	if resultado == 0 {
		for i := 0; i < len(projetos.Items); i++ {
			fmt.Printf("Projeto: %+v\r\n", projetos.Items[i].Metadata.Name)
		}
	} else {
		fmt.Println("[main] Projetos não encontrados em ", url)
	}
}
