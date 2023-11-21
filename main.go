package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {
	// Configurar um endpoint para lidar com notificações
	http.HandleFunc("/webhook", notificationHandler)

	// Iniciar o servidor na porta 8080
	fmt.Println("Servidor ouvindo na porta 8080...")
	http.ListenAndServe(":8080", nil)
}


func notificationHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para processar a notificação recebida do Mercado Livre
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da notificação", http.StatusInternalServerError)
		return
	}

	// Aqui, você pode processar os dados recebidos do Mercado Livre
	fmt.Println("Notificação recebida:", string(body))

	// Responda ao Mercado Livre com um status 200 para indicar que a notificação foi recebida com sucesso
	w.WriteHeader(http.StatusOK)
}




