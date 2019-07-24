package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetRoute recuperar Route
func GetRoute(token string, url string, projeto string, nome string) (resultado int, route model.Route) {
	//token := GetToken(url)
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRoute] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		route = model.Route{}
		err = json.Unmarshal(corpo, &route)
		if err != nil {
			fmt.Println("[GetRoute] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, route
}

// GetRouteString recuperar Route
func GetRouteString(token string, url string, projeto string, nome string) (resultado int, routeString string) {
	//token := GetToken(url)
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRouteString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routeString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, routeString
}

// ListRoute listar todos Routes
func ListRoute(token string, url string) (resultado int, routes model.Routes) {
	//token := GetToken(url)
	endpoint := url + apiRoutes + "routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoute] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routes = model.Routes{}
		err = json.Unmarshal(corpo, &routes)
		if err != nil {
			fmt.Println("[ListRoute] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, routes
}

// ListRouteString listar todos Routes
func ListRouteString(token string, url string) (resultado int, routesString string) {
	//token := GetToken(url)
	endpoint := url + apiRoutes + "routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRouteString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, routesString
}

// ListRouteProjeto listar Routes por projetos
func ListRouteProjeto(token string, url string, projeto string) (resultado int, routes model.Routes) {
	//token := GetToken(url)
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRouteProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routes = model.Routes{}
		err = json.Unmarshal(corpo, &routes)
		if err != nil {
			fmt.Println("[ListRouteProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, routes
}

// ListRouteProjetoString listar Routes por projetos
func ListRouteProjetoString(token string, url string, projeto string) (resultado int, routesString string) {
	//token := GetToken(url)
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRouteString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, routesString
}
