package main

import (
	"fmt"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)
	resultado, dcs := utils.ListPodProjeto(token, url, "teste-01")
	//resultado, dcs := utils.ListPod(url)
	if resultado == 0 {
		for i := 0; i < len(dcs.Items); i++ {
			fmt.Printf("Pod: %+v\r\n", dcs.Items[i].Metadata.Name)
			fmt.Printf("Deploymentconfig: %+v\r\n", dcs.Items[i].Metadata.Labels.Deploymentconfig)
			//fmt.Printf("Pod: %+v\r\n", dcs.Items[i])
			//image := pods.Items[i].Spec.Template.Spec.Containers[0].Name
			break
		}
	} else {
		fmt.Println("[main] Pods não encontrados em ", url)
	}
}
