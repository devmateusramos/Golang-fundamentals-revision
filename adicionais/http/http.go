package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Acesso get em home")
	w.Write([]byte("Olá mundo!"))
}
func main() {
	// HTTP é um protocolo de comunicação- base da comunicação web
	// etc...
	// URI - identificador do recurso
	// Método - O que fazer com o recurso... get, post...
	http.HandleFunc("/home", home)

	log.Fatalln(http.ListenAndServe(":5000", nil))
}
