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

	fmt.Println("Running on port 5000")
	log.Fatalln(http.ListenAndServe(":5000", router))
}
