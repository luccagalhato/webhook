package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Configurar o endpoint para receber webhooks
	http.HandleFunc("/webhook", handleWebhook)

	// Iniciar o servidor HTTP
	port := 8080
	fmt.Printf("Servidor Webhook iniciado na porta %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// Ler o corpo da requisição
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	// Processar o payload do webhook
	processWebhook(body)

	// Responder ao Mercado Livre para confirmar o recebimento do webhook
	w.WriteHeader(http.StatusOK)
}

func processWebhook(payload []byte) {
	// Aqui você pode implementar a lógica para lidar com o payload do webhook
	// Decodificar o payload e realizar ações com base nos eventos recebidos
	fmt.Println("Webhook recebido:")
	fmt.Println(string(payload))
	// Adicione sua lógica de manipulação de webhook aqui
}
