package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetProjeto recuperar projeto
func GetProjeto(token string, url string, nome string) (resultado int, proj model.Projeto) {
	//token := GetToken(url)
	endpoint := url + apiProjeto + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		proj = model.Projeto{}
		err = json.Unmarshal(corpo, &proj)
		if err != nil {
			fmt.Println("[GetProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, proj
}

// Projetos listar projetos
func Projetos(token string, url string) (resultado int, projetos model.Projetos) {
	//token := GetToken(url)
	endpoint := url + apiProjeto

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[Projetos] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		projetos = model.Projetos{}
		err = json.Unmarshal(corpo, &projetos)
		if err != nil {
			fmt.Println("[Projetos] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, projetos
}

// GetNamespace recuperar namespace
func GetNamespace(token string, url string, nome string) (resultado int, proj model.Projeto) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetNamespace] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		proj = model.Projeto{}
		err = json.Unmarshal(corpo, &proj)
		if err != nil {
			fmt.Println("[GetNamespace] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, proj
}

// Namespaces listar namespaces
func Namespaces(token string, url string) (resultado int, projetos model.Projetos) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[Namespaces] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		projetos = model.Projetos{}
		err = json.Unmarshal(corpo, &projetos)
		if err != nil {
			fmt.Println("[Namespaces] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, projetos
}
