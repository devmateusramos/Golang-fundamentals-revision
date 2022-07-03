package main

import (
	router "api-devbook/src"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API")

	routes := router.Gerar()
	log.Fatalln(http.ListenAndServe(":5000", routes))
}
