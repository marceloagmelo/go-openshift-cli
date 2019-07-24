package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL")
	token := utils.GetToken(url)
	//resultado, is := utils.ListDeploymentConfigProjeto(token, url, "teste-01")
	resultado, is := utils.ListImageStream(token, url)
	if resultado == 0 {
		for i := 0; i < len(is.Items); i++ {
			fmt.Printf("Dc: %+v\r\n", is.Items[i].Metadata.Name)
			//fmt.Printf("Dc: %+v\r\n", dcs)
		}
	} else {
		fmt.Println("[main] Dcs não encontrados em ", url)
	}
}
