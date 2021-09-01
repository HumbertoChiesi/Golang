package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CreateUser).Methods(http.MethodPost)        //"POST"
	router.HandleFunc("/usuarios", server.SearchUsers).Methods(http.MethodGet)        //"GET"
	router.HandleFunc("/usuarios/{id}", server.SearchUser).Methods(http.MethodGet)    //"GET"
	router.HandleFunc("/usuarios/{id}", server.UpdateUser).Methods(http.MethodPut)    //"PUT"
	router.HandleFunc("/usuarios/{id}", server.DeleteUser).Methods(http.MethodDelete) //"Delete"

	fmt.Println("Listening to the server port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
