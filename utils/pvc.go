package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetPvc recuperar PVC
func GetPvc(token string, url string, projeto string, nome string) (resultado int, pvc model.Pvc) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/persistentvolumeclaims/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetPvc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvc = model.Pvc{}
		err = json.Unmarshal(corpo, &pvc)
		if err != nil {
			fmt.Println("[GetPvc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pvc
}

// GetPvcString recuperar PVC
func GetPvcString(token string, url string, projeto string, nome string) (resultado int, pvcString string) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/persistentvolumeclaims/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetPvcString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, pvcString
}

// ListPvc listar todos PVC
func ListPvc(token string, url string) (resultado int, pvcs model.Pvcs) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcs = model.Pvcs{}
		err = json.Unmarshal(corpo, &pvcs)
		if err != nil {
			fmt.Println("[ListPvc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pvcs
}

// ListPvcString listar todos PVC
func ListPvcString(token string, url string) (resultado int, pvcsString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvcString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, pvcsString
}

// ListPvcProjeto listar PVC por projetos
func ListPvcProjeto(token string, url string, projeto string) (resultado int, pvcs model.Pvcs) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvcProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcs = model.Pvcs{}
		err = json.Unmarshal(corpo, &pvcs)
		if err != nil {
			fmt.Println("[ListPvcProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pvcs
}

// ListPvcProjetoString listar PVC por projetos
func ListPvcProjetoString(token string, url string, projeto string) (resultado int, pvcsString string) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvcString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, pvcsString
}
