package main

import (
	"fmt"
	"strings"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
        token := utils.GetToken(url)

	resultado, dcs := utils.ListDc(token, url)
	if resultado == 0 {
		for i := 0; i < len(dcs.Items); i++ {
			//nomeDc := dcs.Items[i].Metadata.Name
			//nomeProjeto := dcs.Items[i].Metadata.Namespace
			image := dcs.Items[i].Spec.Template.Spec.Containers[0].Image
			aImage := strings.Split(image, "/")
			//fmt.Printf("Projeto: %+v\r\n", nomeProjeto)
			//fmt.Printf("Dc: %+v\r\n", nomeDc)
			fmt.Printf("DC image: %+v\r\n", image)
			for j := 0; j < len(aImage); j++ {
				fmt.Printf("DC image array [%d]: %+v\r\n", j, aImage[j])
			}
			//fmt.Printf("DC image array: %+v\r\n", aImage[len(aImage)-1])
		}
	} else {
		fmt.Println("[main] Dcs não encontrados")
	}
}
