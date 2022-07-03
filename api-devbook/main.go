package main

import (
	"api-devbook/src/config"
	"api-devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	routes := router.Gerar()
	fmt.Printf("Server is running on port %d", config.Porta)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), routes))
}
