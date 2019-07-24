package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL") //utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, daemonset := utils.GetDaemonSetString(token, url, "kube-service-catalog", "apiserver")
	if resultado > 0 {
		fmt.Println("[main] O daemonset debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", daemonset)
}
