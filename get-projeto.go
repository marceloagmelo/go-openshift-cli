package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL") //utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, proj := utils.GetNamespace(token, url, "teste-01")
	if resultado > 0 {
		fmt.Println("[main] O projeto solr-lfr-pre não encontrado em ", url)
		return
	}
	fmt.Printf("Post é: %+v\r\n", proj)
}
