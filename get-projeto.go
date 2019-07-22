package main

import (
	"fmt"

	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, proj := utils.GetNamespace(token, url, "acb-dev")
	if resultado > 0 {
		fmt.Println("[main] O projeto solr-lfr-pre não encontrado em ", url)
		return
	}
	fmt.Printf("Post é: %+v\r\n", proj)
}
