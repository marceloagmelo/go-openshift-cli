package main

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL")
	token := utils.GetToken(url)

	resultado, projetos := utils.Projetos(token, url)
	if resultado == 0 {
		for i := 0; i < len(projetos.Items); i++ {
			projeto := projetos.Items[i].Metadata.Name
			fmt.Printf("Projeto: %+v\r\n", projeto)

			resultado, pods := utils.ListPodProjeto(token, url, projeto)
			if resultado == 0 {
				for i := 0; i < len(pods.Items); i++ {
					nomePod := pods.Items[i].Metadata.Name
					image := pods.Items[i].Spec.Containers[0].Image
					fmt.Printf("Nome: %+v\r\n", nomePod)
					fmt.Printf("Image: %+v\r\n", image)
				}
			} else {
				fmt.Println("[main] Dcs não encontrados no projeto ", projeto)
			}
		}
	} else {
		fmt.Println("[main] Projetos não encontrados em ", url)
	}
}
