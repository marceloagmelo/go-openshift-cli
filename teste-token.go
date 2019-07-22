package main

import (
	"fmt"

	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("dev")
	token := utils.GetToken(url)
	fmt.Println("TOKEN == ", token)
	for i := 0; i < 5000; i++ {
		resultado, dc := utils.GetDcTeste(token, url, "solr-lfr-pre", "debug")
		if resultado == 0 {
			fmt.Println("[dc] ", dc.Metadata.Name)
		} else {
			fmt.Println("[ERRO] ")
		}
	}
}
