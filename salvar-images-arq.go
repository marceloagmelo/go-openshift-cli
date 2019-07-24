package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marceloagmelo/go-openshift-cli/utils"
)

func main() {
	url := os.Getenv("OPENSHIFT_URL")
	token := utils.GetToken(url)

	resultado, dcs := utils.ListDc(token, url)
	if resultado == 0 {
		arquivoTXT, err := os.Create("/home/marcelo/imagens.txt")
		if err != nil {
			fmt.Println("[main] Houve um erro ao criar o arquivo TXT. Erro: ", err.Error())
			return
		}
		defer arquivoTXT.Close()
		escritor := bufio.NewWriter(arquivoTXT)

		for i := 0; i < len(dcs.Items); i++ {
			nomeProjeto := dcs.Items[i].Metadata.Namespace
			nomeDc := dcs.Items[i].Metadata.Name
			image := dcs.Items[i].Spec.Template.Spec.Containers[0].Image
			aImage := strings.Split(image, "/")
			escritor.WriteString("\r\n")
			escritor.WriteString("DC image: " + image)
			if len(aImage) > 2 {
				imageName := aImage[1] + "/" + aImage[2]
				escritor.WriteString("\r\n")
				escritor.WriteString("DC image: " + imageName)

				resultado, dc := utils.GetDc(token, url, nomeProjeto, nomeDc)
				if resultado == 0 {
					fmt.Println("[dc] ", dc.Metadata.Name)
				} else {
					fmt.Printf("[recuperarDc] Dc %s não encontrado no projeto %s ambiente %s", nomeDc, nomeProjeto, url)
				}
			}
		}
		escritor.Flush()
	} else {
		fmt.Println("[main] Dcs não encontrados")
	}
}
