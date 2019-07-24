package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetSecret recuperar Secret
func GetSecret(token string, url string, projeto string, nome string) (resultado int, secret model.Secret) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/secrets/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetSecret] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		secret = model.Secret{}
		err = json.Unmarshal(corpo, &secret)
		if err != nil {
			fmt.Println("[GetSecret] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, secret
}

// GetSecretString recuperar Secret
func GetSecretString(token string, url string, projeto string, nome string) (resultado int, secretString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/secrets/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetSecretString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		secretString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, secretString
}

// ListSecret listar todos Secret
func ListSecret(token string, url string) (resultado int, secrets model.Secrets) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "secrets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListSecret] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		secrets = model.Secrets{}
		err = json.Unmarshal(corpo, &secrets)
		if err != nil {
			fmt.Println("[ListSecret] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, secrets
}

// ListSecretString listar todos Secret
func ListSecretString(token string, url string) (resultado int, secretsString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "secrets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListSecretString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		secretsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, secretsString
}

// ListSecretProjeto listar Secret por projetos
func ListSecretProjeto(token string, url string, projeto string) (resultado int, secrets model.Secrets) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/secrets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListSecretProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		secrets = model.Secrets{}
		err = json.Unmarshal(corpo, &secrets)
		if err != nil {
			fmt.Println("[ListSecretProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, secrets
}

// ListSecretProjetoString listar Secret por projetos
func ListSecretProjetoString(token string, url string, projeto string) (resultado int, secretsString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "namespaces/" + projeto + "/secrets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListSecretString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		secretsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, secretsString
}
