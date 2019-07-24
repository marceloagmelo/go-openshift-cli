package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL") //utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, pvc := utils.GetPvcString(token, url, "teste-03", "rabbitmq-cluster-storage")
	if resultado > 0 {
		fmt.Println("[main] O pvc debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", pvc)
}
