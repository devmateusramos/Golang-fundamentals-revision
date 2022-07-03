package main

import (
	"crud-basic/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).
		Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).
		Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).
		Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).
		Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).
		Methods(http.MethodDelete)

	fmt.Println("Running on port 5000")
	log.Fatalln(http.ListenAndServe(":5000", router))
}
