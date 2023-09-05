package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CEP struct {
	Cep         string `json:"cep"`
	AddressType string `json:"address_type"`
	AddressName string `json:"address_name"`
	Address     string `json:"address"`
	State       string `json:"state"`
	District    string `json:"district"`
	Lat         string `json:"lat"`
	Lng         string `json:"lng"`
	City        string `json:"city"`
	CityIBGE    string `json:"city_ibge"`
	DDD         string `json:"ddd"`
}

func main() {
	http.HandleFunc("/cep/", cepHandler)

	fmt.Println("Servidor disponível em http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}

func cepHandler(w http.ResponseWriter, r *http.Request) {
	cep := strings.TrimPrefix(r.URL.Path, "/cep/")
	if len(cep) != 8 {
		errorResponse(w, "CEP inválido", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://cep.awesomeapi.com.br/json/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Erro ao fazer a requisição:", err)
		errorResponse(w, "Erro ao buscar o CEP", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorResponse(w, "CEP não encontrado", http.StatusNotFound)
		return
	}

	var cepData CEP
	err = json.NewDecoder(resp.Body).Decode(&cepData)
	if err != nil {
		log.Println("Erro ao decodificar a resposta:", err)
		errorResponse(w, "Erro ao buscar o CEP", http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, cepData)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func errorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]string{"error": message}
	w.WriteHeader(statusCode)
	writeJSONResponse(w, response)
}
