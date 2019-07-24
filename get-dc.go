package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)

	resultado, dc := utils.GetDeploymentConfigString(token, url, "teste-01", "nodemysql")
	if resultado > 0 {
		fmt.Println("[main] O dc debug não encontrado em ", url)
		return
	}
	fmt.Printf("%+v\r\n", dc)
	//fmt.Printf("DC é: %+v\r\n", dc)
	//image := utils.GetDcImage(dc)
	//fmt.Printf("DC image: %+v\r\n", image)
}
