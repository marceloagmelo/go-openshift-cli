package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := utils.URLGen("pre")
	token := utils.GetToken(url)
	//resultado, dcs := utils.ListDeploymentConfigProjeto(token, url, "teste-01")
	resultado, dcs := utils.ListDeploymentConfig(token, url)
	os.Mkdir("/home/marcelo/jsons", 0700)
	os.Mkdir("/home/marcelo/jsons/teste", 0700)
	if resultado == 0 {
		for i := 0; i < len(dcs.Items); i++ {
			//arquivo := "/home/marcelo/jsons/teste/" + dcs.Items[i].Metadata.Name + ".json"
			//arquivo := "/home/marcelo/jsons/teste/teste.json"
			fmt.Printf("Dc: %+v\r\n", dcs.Items[i].Metadata.Name)
			//fmt.Printf("Dc: %+v\r\n", dcs.Items[i])
			//salvarAquivo(arquivo, fmt.Sprintf("%+v", dcs.Items[i]))
			//lerDc(token, dcs.Items[i].Metadata.Namespace, dcs.Items[i].Metadata.Name)
		}
		//salvarAquivo(arquivo, fmt.Sprintf("%+v", dcs))

	} else {
		fmt.Println("[main] Dcs não encontrados em ", url)
	}
}

func salvarAquivo(arquivo string, texto string) {
	arquivoJSON, err := os.Create(arquivo)
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
	}
	defer arquivoJSON.Close()
	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString(texto)
	escritor.Flush()
}
