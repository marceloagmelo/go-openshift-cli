package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, service := utils.GetServiceString(token, url, "teste-01", "nodemysql")
	if resultado > 0 {
		fmt.Println("[main] O service debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", service)
}
