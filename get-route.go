package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, route := utils.GetRouteString(token, url, "teste-01", "vs-nodemysql")
	if resultado > 0 {
		fmt.Println("[main] O route debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", route)
}
