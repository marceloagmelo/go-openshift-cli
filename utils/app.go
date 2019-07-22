package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetDc recuperar DC
func GetDc(token string, url string, projeto string, nome string) (resultado int, dc model.Dc) {
	//token := GetToken(url)
	endpoint := url + apiNamespaceApps + projeto + "/deploymentconfigs/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dc = model.Dc{}
		err = json.Unmarshal(corpo, &dc)
		if err != nil {
			fmt.Println("[GetDc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dc
}

// GetDcString recuperar DC
func GetDcString(token string, url string, projeto string, nome string) (resultado int, dcString string) {
	//token := GetToken(url)
	endpoint := url + apiNamespaceApps + projeto + "/deploymentconfigs/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcString
}

// ListDc listar todos deploymentconfig
func ListDc(token string, url string) (resultado int, dcs model.Dcs) {
	//token := GetToken(url)
	endpoint := url + apiApps + "deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcs = model.Dcs{}
		err = json.Unmarshal(corpo, &dcs)
		if err != nil {
			fmt.Println("[ListDc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dcs
}

// ListDcString listar todos deploymentconfig
func ListDcString(token string, url string) (resultado int, dcsString string) {
	//token := GetToken(url)
	endpoint := url + apiApps + "deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcsString
}

// ListDcProjeto listar deploymentconfig por projetos
func ListDcProjeto(token string, url string, projeto string) (resultado int, dcs model.Dcs) {
	//token := GetToken(url)
	endpoint := url + apiNamespaceApps + projeto + "/deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDcProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcs = model.Dcs{}
		err = json.Unmarshal(corpo, &dcs)
		if err != nil {
			fmt.Println("[ListDcProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dcs
}

// GetDcImage recuperar DC
func GetDcImage(dc model.Dc) string {
	return dc.Spec.Template.Spec.Containers[0].Image
}
