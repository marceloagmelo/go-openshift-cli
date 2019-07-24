package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL")
	token := utils.GetToken(url)
	resultado, secrets := utils.ListSecretProjeto(token, url, "teste-03")
	//resultado, secrets := utils.ListSecret(token, url)
	if resultado == 0 {
		for i := 0; i < len(secrets.Items); i++ {
			fmt.Printf("Secret %+v\r\n", secrets.Items[i].Metadata.Name)
		}
	} else {
		fmt.Println("[main] secrets não encontrados em ", url)
	}
}
