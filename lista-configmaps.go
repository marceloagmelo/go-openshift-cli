package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL") //utils.URLGen("pre")
	token := utils.GetToken(url)
	//resultado, configmaps := utils.ListConfigMapProjeto(token, url, "teste-04")
	resultado, configmaps := utils.ListConfigMap(token, url)
	if resultado == 0 {
		for i := 0; i < len(configmaps.Items); i++ {
			fmt.Printf("ConfigMap %+v\r\n", configmaps.Items[i].Metadata.Name)
		}
	} else {
		fmt.Println("[main] configmaps não encontrados em ", url)
	}
}
