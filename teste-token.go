package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL")
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
