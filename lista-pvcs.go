package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)
	resultado, pvcs := utils.ListPvcProjeto(token, url, "teste-03")
	//resultado, pvcs := utils.ListPvc(token, url)
	if resultado == 0 {
		for i := 0; i < len(pvcs.Items); i++ {
			fmt.Printf("PVC %+v\r\n", pvcs.Items[i].Metadata.Name)
		}
	} else {
		fmt.Println("[main] Routes não encontrados em ", url)
	}
}
