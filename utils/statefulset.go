package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetStateFulSet recuperar StateFulSet
func GetStateFulSet(token string, url string, projeto string, nome string) (resultado int, statefulset model.StateFulSet) {
	//token := GetToken(url)
	endpoint := url + apisAppsv1beta1Namespace + projeto + "/statefulsets/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetStateFulSet] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		statefulset = model.StateFulSet{}
		err = json.Unmarshal(corpo, &statefulset)
		if err != nil {
			fmt.Println("[GetStateFulSet] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, statefulset
}

// GetStateFulSetString recuperar StateFulSet
func GetStateFulSetString(token string, url string, projeto string, nome string) (resultado int, statefulsetString string) {
	//token := GetToken(url)
	endpoint := url + apisAppsv1beta1Namespace + projeto + "/statefulsets/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetStateFulSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		statefulsetString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, statefulsetString
}

// ListStateFulSet listar todos StateFulSet
func ListStateFulSet(token string, url string) (resultado int, statefulsets model.StateFulSets) {
	//token := GetToken(url)
	endpoint := url + apisAppsv1beta1 + "statefulsets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListStateFulSet] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		statefulsets = model.StateFulSets{}
		err = json.Unmarshal(corpo, &statefulsets)
		if err != nil {
			fmt.Println("[ListStateFulSet] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, statefulsets
}

// ListStateFulSetString listar todos StateFulSet
func ListStateFulSetString(token string, url string) (resultado int, statefulsetsString string) {
	//token := GetToken(url)
	endpoint := url + apisAppsv1beta1 + "statefulsets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListStateFulSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		statefulsetsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, statefulsetsString
}

// ListStateFulSetProjeto listar StateFulSet por projetos
func ListStateFulSetProjeto(token string, url string, projeto string) (resultado int, statefulsets model.StateFulSets) {
	//token := GetToken(url)
	endpoint := url + apisAppsv1beta1Namespace + projeto + "/statefulsets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListStateFulSetProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		statefulsets = model.StateFulSets{}
		err = json.Unmarshal(corpo, &statefulsets)
		if err != nil {
			fmt.Println("[ListStateFulSetProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, statefulsets
}

// ListStateFulSetProjetoString listar StateFulSet por projetos
func ListStateFulSetProjetoString(token string, url string, projeto string) (resultado int, statefulsetsString string) {
	//token := GetToken(url)
	endpoint := url + apisAppsv1beta1Namespace + projeto + "/statefulsets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListStateFulSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		statefulsetsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, statefulsetsString
}
