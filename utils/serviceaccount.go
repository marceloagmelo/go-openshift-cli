package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// getServiceAccount recuperar ServiceAccount
func getServiceAccount(token string, url string, projeto string, nome string) (resultado int, serviceAccount model.ServiceAccount) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getServiceAccount] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccount = model.ServiceAccount{}
		err = json.Unmarshal(corpo, &serviceAccount)
		if err != nil {
			fmt.Println("[getServiceAccount] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, serviceAccount
}

// getServiceAccountString recuperar ServiceAccount
func getServiceAccountString(token string, url string, projeto string, nome string) (resultado int, serviceAccountString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getServiceAccountString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccountString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceAccountString
}

// ListServiceAccount listar todos serviceaccounts
func ListServiceAccount(token string, url string) (resultado int, serviceAccount model.ServiceAccount) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceAccount] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccount = model.ServiceAccount{}
		err = json.Unmarshal(corpo, &serviceAccount)
		if err != nil {
			fmt.Println("[ListServiceAccount] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, serviceAccount
}

// ListserviceAccountString listar todos serviceaccounts
func ListserviceAccountString(token string, url string) (resultado int, serviceAccountString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListserviceAccountString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccountString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceAccountString
}

// ListServiceAccountProjeto listar serviceaccounts por projetos
func ListServiceAccountProjeto(token string, url string, projeto string) (resultado int, serviceAccount model.ServiceAccount) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceAccountProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccount = model.ServiceAccount{}
		err = json.Unmarshal(corpo, &serviceAccount)
		if err != nil {
			fmt.Println("[ListServiceAccountProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, serviceAccount
}

// ListServiceAccountProjetoString listar serviceaccounts por projetos
func ListServiceAccountProjetoString(token string, url string, projeto string) (resultado int, serviceAccountString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListserviceAccountString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccountString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceAccountString
}
