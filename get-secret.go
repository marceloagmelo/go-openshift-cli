package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, secret := utils.GetSecretString(token, url, "teste-03", "rabbitmq-cluster-secret")
	if resultado > 0 {
		fmt.Println("[main] O secret debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", secret)
}