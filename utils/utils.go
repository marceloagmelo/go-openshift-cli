package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/marceloagmelo/go-openshift-cli/model"
	yaml "gopkg.in/yaml.v2"
)

//var configfile = "/home/marcelo/go/src/gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/conf/config.yaml"
var apiProjeto = "/apis/project.openshift.io/v1/projects/"
var apiApps = "/apis/apps.openshift.io/v1/"
var apiPods = "/api/v1/pods/"
var apiV1 = "/api/v1/"
var apiRoutes = "/apis/route.openshift.io/v1/"
var apisAppsv1beta1 = "/apis/apps/v1beta1/"
var apisImageV1 = "/apis/image.openshift.io/v1"
var apisAuthorizationOpenshiftV1 = "/apis/authorization.openshift.io/v1/"
var apisExtensionsV1beta1 = "/apis/extensions/v1beta1/"
var apiBuilds = "/apis/build.openshift.io/v1/"

var data = `
scheme: https
ambientePre: console.openshift-311.lab
ambientePro: console.openshift-311.lab
username: admin
password: admin123
`

type generalConfig struct {
	Scheme      string `yaml:"scheme"`
	AmbientePre string `yaml:"ambientePre"`
	AmbientePro string `yaml:"ambientePro"`
	Usuario     string `yaml:"username"`
	Senha       string `yaml:"password"`
}

// generalConfigLoad load general configuration from conf/config.yaml
func generalConfigLoad() (*generalConfig, error) {
	var config generalConfig

	/*dataBytes, err := ioutil.ReadFile(configfile)
	if err != nil {
		return nil, err
	}*/
	//err = yaml.Unmarshal([]byte(dataBytes), &config)
	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// URLGen generates target URL.
func URLGen(ambiente string) string {
	config, err := generalConfigLoad()
	if err != nil {
		fmt.Println("URLGen:", err)
		os.Exit(1)
	}
	url := config.Scheme + "://" + config.AmbientePre + ":8443"
	if ambiente == "pro" {
		url = config.Scheme + "://" + config.AmbientePro + ":8443"
	}
	return url
}

// GetToken recuperar Token do usuário.
func GetToken(url string) string {
	config, err := generalConfigLoad()
	if err != nil {
		fmt.Println("URLGen:", err)
		os.Exit(1)
	}

	endpoint := url + "/oauth/authorize?client_id=openshift-challenging-client&response_type=token"

	cmdCurl := "curl -s -u " + config.Usuario + ":" + config.Senha + " -kI '" + endpoint + "' | grep -oP 'access_token=\\K[^&]*'"

	resultado, resposta := ExecCurl(cmdCurl)

	if resultado > 0 {
		fmt.Println("[GetToken] Erro ao executar o CURL.", err.Error())
	}
	return resposta
}

//GetUsuarioSenha
func GetUsuarioSenha() (usuario string, senha string) {
	config, err := generalConfigLoad()
	if err != nil {
		fmt.Println("URLGen:", err)
		os.Exit(1)
	}

	usuario = config.Usuario
	senha = config.Senha

	return usuario, senha
}

// ExecCurl execuctar CURL.
func ExecCurl(strCurl string) (resultado int, resposta string) {
	resultado = 0
	cmd := exec.Command("/bin/bash", "-c", strCurl)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ExecCurl:", err)
		resultado = 1
	} else {
		resposta = string(out)
	}

	return resultado, resposta
}

// GetRequest recuperar a requisição
func GetRequest(token string, endpoint string) (resultado int, resposta *http.Response) {
	resultado = 0
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 20,
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("[GetRequest] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[GetRequest] Erro ao abrir a pagina. Erro: ", err.Error())
		resultado = 1
	}
	return resultado, resposta
}

// GetRequestJSON recuperar a requisição
func GetRequestJSON(token string, endpoint string) (resultado int, dc model.Dc) {
	resultado = 0

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 90,
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("[GetRequestJSON] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[GetRequestJSON] Erro ao abrir a pagina. Erro: ", err.Error())
		resultado = 1
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRequestJSON] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		//dc := model.Dc{}
		err = json.Unmarshal(corpo, &dc)
		if err != nil {
			fmt.Println("[GetDc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}

	return resultado, dc
}

// GetRequestString recuperar a requisição
func GetRequestString(token string, endpoint string) (resultado int, dcString string) {
	resultado = 0

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 90,
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("[GetRequestString] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[GetRequestString] Erro ao abrir a pagina. Erro: ", err.Error())
		resultado = 1
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRequestString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcString
}
