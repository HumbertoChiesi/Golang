package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!"))
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Load users page!!"))
}

func main() {
	//HTTP é um protocolo de comunicação
	//Cliente (Faz requisição) - Servidor (Processa requisição e envia resposta)

	//Request - Response
	//Rotas
	//URI - Identificador do recurso
	//Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", homePage)

	http.HandleFunc("/users", usersPage)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
