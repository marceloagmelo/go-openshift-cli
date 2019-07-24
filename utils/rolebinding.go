package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetRoleBinding recuperar RoleBinding
func GetRoleBinding(token string, url string, projeto string, nome string) (resultado int, roleBinding model.RoleBinding) {
	//token := GetToken(url)
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/rolebindings/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRoleBinding] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		roleBinding = model.RoleBinding{}
		err = json.Unmarshal(corpo, &roleBinding)
		if err != nil {
			fmt.Println("[GetRoleBinding] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, roleBinding
}

// GetRoleBindingString recuperar RoleBinding
func GetRoleBindingString(token string, url string, projeto string, nome string) (resultado int, roleBindingString string) {
	//token := GetToken(url)
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/rolebindings/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRoleBindingString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		roleBindingString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, roleBindingString
}

// ListRoleBinding listar todos RoleBindings
func ListRoleBinding(token string, url string) (resultado int, rolebinding model.RoleBindings) {
	//token := GetToken(url)
	endpoint := url + apisAuthorizationOpenshiftV1 + "rolebindings"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoleBinding] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		rolebinding = model.RoleBindings{}
		err = json.Unmarshal(corpo, &rolebinding)
		if err != nil {
			fmt.Println("[ListRoleBinding] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, rolebinding
}

// ListRoleBindingString listar todos RoleBindings
func ListRoleBindingString(token string, url string) (resultado int, rolebindingString string) {
	//token := GetToken(url)
	endpoint := url + apisAuthorizationOpenshiftV1 + "rolebindings"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoleBindingString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		rolebindingString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, rolebindingString
}

// ListRoleBindingProjeto listar RoleBindings por projetos
func ListRoleBindingProjeto(token string, url string, projeto string) (resultado int, rolebinding model.RoleBindings) {
	//token := GetToken(url)
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/rolebindings"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoleBindingProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		rolebinding = model.RoleBindings{}
		err = json.Unmarshal(corpo, &rolebinding)
		if err != nil {
			fmt.Println("[ListRoleBindingProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, rolebinding
}

// ListRoleBindingProjetoString listar RoleBindings por projetos
func ListRoleBindingProjetoString(token string, url string, projeto string) (resultado int, rolebindingString string) {
	//token := GetToken(url)
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/rolebindings"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoleBindingString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		rolebindingString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, rolebindingString
}
