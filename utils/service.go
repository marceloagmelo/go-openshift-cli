package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetService recuperar SVC
func GetService(token string, url string, projeto string, nome string) (resultado int, service model.Service) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetService] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		service = model.Service{}
		err = json.Unmarshal(corpo, &service)
		if err != nil {
			fmt.Println("[GetService] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, service
}

// GetServiceString recuperar SVC
func GetServiceString(token string, url string, projeto string, nome string) (resultado int, serviceString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetServiceString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceString
}

// ListService listar todos SVC
func ListService(token string, url string) (resultado int, services model.Services) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListService] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		services = model.Services{}
		err = json.Unmarshal(corpo, &services)
		if err != nil {
			fmt.Println("[ListService] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, services
}

// ListServiceString listar todos SVC
func ListServiceString(token string, url string) (resultado int, servicesString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		servicesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, servicesString
}

// ListServiceProjeto listar SVC por projetos
func ListServiceProjeto(token string, url string, projeto string) (resultado int, services model.Services) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		services = model.Services{}
		err = json.Unmarshal(corpo, &services)
		if err != nil {
			fmt.Println("[ListServiceProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, services
}

// ListServiceProjetoString listar SVC por projetos
func ListServiceProjetoString(token string, url string, projeto string) (resultado int, servicesString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		servicesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, servicesString
}
