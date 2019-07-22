package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)
	resultado, statefulset := utils.ListStateFulSetProjeto(token, url, "teste-03")
	//resultado, statefulset := utils.ListStateFulSet(token, url)
	if resultado == 0 {
		for i := 0; i < len(statefulset.Items); i++ {
			fmt.Printf("Route %+v\r\n", statefulset.Items[i].Metadata.Name)
		}
	} else {
		fmt.Println("[main] statefulset não encontrados em ", url)
	}
}
