package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL") //utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, route := utils.GetRouteString(token, url, "teste-01", "vs-nodemysql")
	if resultado > 0 {
		fmt.Println("[main] O route debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", route)
}
