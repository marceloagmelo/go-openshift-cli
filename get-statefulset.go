package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, statefulset := utils.GetStateFulSetString(token, url, "teste-03", "rabbitmq-cluster")
	if resultado > 0 {
		fmt.Println("[main] O statefulset debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", statefulset)
}
