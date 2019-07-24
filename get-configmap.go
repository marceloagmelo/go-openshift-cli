package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := "https://console.openshift-311.lab:8443"
	token := utils.GetToken(url, "admin", "admin123")

	resultado, configmap := utils.GetConfigMapString(token, url, "teste-04", "redis-conf")
	if resultado > 0 {
		fmt.Println("[main] O configmap debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", configmap)
}
