package main

import (
	"fmt"
	"log"
	"errors"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)

const base_url = "http://viacep.com.br/ws/"
const serverport = "3000"

type CepResult struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Unidade string `json:"unidade"`
	Ibge string `json:"ibge"`
	Gia string `json:"gia"`
}

func getCep(id string) (*CepResult, error) {

	url := base_url + id + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Erro de acesso a URL")
	}
	
	defer resp.Body.Close()

	if (resp.StatusCode != 200) {
		return nil, errors.New("Status code != 200")
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Erro na leitura HTTP")
	}

	var result CepResult	
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, errors.New("JSON Parser error")
	}

	return &result, nil
}
	
func MainHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := getCep("01311200")
	if (resp != nil)  {
		fmt.Println(resp.Logradouro)
		fmt.Println(resp.Complemento)
		fmt.Println(resp.Bairro)
		fmt.Println(resp.Localidade)
		fmt.Println(resp.Uf)
		fmt.Println(resp.Unidade)
		fmt.Println(resp.Ibge)
		fmt.Println(resp.Gia)
	} else {
		fmt.Println(err)
	}
}

func openWebBrowser() {
	http.Get("http://localhost:" + serverport)	
}

func main() {
	go openWebBrowser()
	router := mux.NewRouter()
	router.HandleFunc("/", MainHandler)
	log.Fatal(http.ListenAndServe(":" + serverport, router))
}
