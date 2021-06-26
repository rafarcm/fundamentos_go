package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página de usuários!"))
}

func main() {
	// HTTP é um protocolo de comunicação - Base de comunicação web
	// Cliente (Faz requesição) - Servidor (Processa requisição e envia resposta)
	// Request - Response
	// Rotas: URI (Identificador do recurso) - Métodos - GET, POST, PUT, DELETE

	fmt.Println("HTTP")
	fmt.Println("-----------------------------------------------")

	// Recebe a URI e a função que recebe a requisição e processa ela
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil)) // Sobe um servidor rodando na porta 5000
}
