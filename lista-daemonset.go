package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL")
	token := utils.GetToken(url)
	//resultado, daemonset := utils.ListDaemonSetProjeto(token, url, "teste-03")
	resultado, daemonset := utils.ListDaemonSet(token, url)
	if resultado == 0 {
		for i := 0; i < len(daemonset.Items); i++ {
			fmt.Printf("Route %+v\r\n", daemonset.Items[i].Metadata.Name)
		}
	} else {
		fmt.Println("[main] daemonset não encontrados em ", url)
	}
}
