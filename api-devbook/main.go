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
	fmt.Println(config.Porta)

	routes := router.Gerar()
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), routes))
}
