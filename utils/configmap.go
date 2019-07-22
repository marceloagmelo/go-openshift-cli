package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetConfigMap recuperar ConfigMap
func GetConfigMap(token string, url string, projeto string, nome string) (resultado int, configMap model.ConfigMap) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/configmaps/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetConfigMap] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configMap = model.ConfigMap{}
		err = json.Unmarshal(corpo, &configMap)
		if err != nil {
			fmt.Println("[GetConfigMap] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, configMap
}

// GetConfigMapString recuperar ConfigMap
func GetConfigMapString(token string, url string, projeto string, nome string) (resultado int, configMapString string) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/configmaps/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetConfigMapString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configMapString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, configMapString
}

// ListConfigMap listar todos ConfigMaps
func ListConfigMap(token string, url string) (resultado int, configmap model.ConfigMaps) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMap] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configmap = model.ConfigMaps{}
		err = json.Unmarshal(corpo, &configmap)
		if err != nil {
			fmt.Println("[ListConfigMap] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, configmap
}

// ListConfigMapString listar todos ConfigMaps
func ListConfigMapString(token string, url string) (resultado int, configmapString string) {
	//token := GetToken(url)
	endpoint := url + apiV1 + "configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMapString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configmapString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, configmapString
}

// ListConfigMapProjeto listar ConfigMaps por projetos
func ListConfigMapProjeto(token string, url string, projeto string) (resultado int, configmap model.ConfigMaps) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMapProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configmap = model.ConfigMaps{}
		err = json.Unmarshal(corpo, &configmap)
		if err != nil {
			fmt.Println("[ListConfigMapProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, configmap
}

// ListConfigMapProjetoString listar ConfigMaps por projetos
func ListConfigMapProjetoString(token string, url string, projeto string) (resultado int, configmapString string) {
	//token := GetToken(url)
	endpoint := url + apiNamespaces + projeto + "/configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMapString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configmapString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, configmapString
}
