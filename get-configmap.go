package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, configmap := utils.GetConfigMapString(token, url, "teste-04", "redis-conf")
	if resultado > 0 {
		fmt.Println("[main] O configmap debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", configmap)
}
