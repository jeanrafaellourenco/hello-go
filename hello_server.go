package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Endpoint struct {
	Path        string   `json:"path"`
	Method      string   `json:"method"`
	Description string   `json:"description"`
	Params      []string `json:"params,omitempty"`
}

type APIEndpoints struct {
	Endpoints []Endpoint `json:"endpoints"`
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/healthcheck", healthHandler)
	http.HandleFunc("/docs", docsHandler)

	fmt.Println("Servidor disponível em http://127.0.0.1:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Olá, Mundo!"}
	writeJSONResponse(w, response)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		name := r.URL.Query().Get("name")
		if name != "" {
			response := map[string]string{"message": fmt.Sprintf("Olá, %s!", name)}
			writeJSONResponse(w, response)
		} else {
			response := map[string]string{"message": "Olá!"}
			writeJSONResponse(w, response)
		}
	} else {
		errorResponse(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		message := r.FormValue("message")
		if message != "" {
			response := map[string]string{"message": fmt.Sprintf("Mensagem recebida: %s", message)}
			writeJSONResponse(w, response)
		} else {
			errorResponse(w, "Parâmetro 'message' ausente", http.StatusBadRequest)
		}
	} else {
		errorResponse(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "API está online e saudável!"}
	writeJSONResponse(w, response)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	endpoints := APIEndpoints{
		Endpoints: []Endpoint{
			{
				Path:        "/",
				Method:      "GET",
				Description: "Endpoint principal que retorna 'Olá, Mundo!'",
			},
			{
				Path:        "/greet",
				Method:      "GET",
				Description: "Saudação personalizada com o parâmetro 'name' na URL",
				Params:      []string{"name"},
			},
			{
				Path:        "/message",
				Method:      "POST",
				Description: "Recebe uma mensagem através de um campo de formulário chamado 'message'",
			},
			{
				Path:        "/healthcheck",
				Method:      "GET",
				Description: "Verifica o status de saúde da API",
			},
		},
	}

	writeJSONResponse(w, endpoints)
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
