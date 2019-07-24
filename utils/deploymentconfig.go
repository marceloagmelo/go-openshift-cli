package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetDeploymentConfig recuperar DC
func GetDeploymentConfig(token string, url string, projeto string, nome string) (resultado int, dc model.DeploymentConfig) {
	//token := GetToken(url)
	endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dc = model.Dc{}
		err = json.Unmarshal(corpo, &dc)
		if err != nil {
			fmt.Println("[GetDeploymentConfig] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dc
}

// GetDeploymentConfigString recuperar DC
func GetDeploymentConfigString(token string, url string, projeto string, nome string) (resultado int, dcString string) {
	//token := GetToken(url)
	endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs/" + nome

	fmt.Println("[endpoint] Erro ao ler o conteudo da pagina. Erro: ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcString
}

// ListDeploymentConfig listar todos deploymentconfig
func ListDeploymentConfig(token string, url string) (resultado int, dcs model.DeploymentConfigs) {
	//token := GetToken(url)
	endpoint := url + apiApps + "deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcs = model.Dcs{}
		err = json.Unmarshal(corpo, &dcs)
		if err != nil {
			fmt.Println("[ListDeploymentConfig] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dcs
}

// ListDeploymentConfigString listar todos deploymentconfig
func ListDeploymentConfigString(token string, url string) (resultado int, dcsString string) {
	//token := GetToken(url)
	endpoint := url + apiApps + "deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcsString
}

// ListDeploymentConfigProjeto listar deploymentconfig por projetos
func ListDeploymentConfigProjeto(token string, url string, projeto string) (resultado int, dcs model.DeploymentConfigs) {
	//token := GetToken(url)
	endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentConfigProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcs = model.DeploymentConfigs{}
		err = json.Unmarshal(corpo, &dcs)
		if err != nil {
			fmt.Println("[ListDeploymentConfigProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dcs
}

// GetDeploymentConfigImage recuperar DC
func GetDeploymentConfigImage(dc model.DeploymentConfig) string {
	return dc.Spec.Template.Spec.Containers[0].Image
}
