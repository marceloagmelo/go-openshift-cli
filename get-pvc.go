package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, pvc := utils.GetPvcString(token, url, "teste-03", "rabbitmq-cluster-storage")
	if resultado > 0 {
		fmt.Println("[main] O pvc debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", pvc)
}
