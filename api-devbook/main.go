package main

import (
	"api-devbook/src/config"
	"api-devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

//func init() {
//	// Para gerar o secret a ser posto no .env
//	chave := make([]byte, 64)
//
//	if _, err := rand.Read(chave); err != nil {
//		log.Fatalln(err)
//	}
//
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//
//	fmt.Println(stringBase64)
//}

func main() {
	config.Carregar()

	routes := router.Gerar()
	fmt.Printf("Server is running on port %d \n", config.Porta)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), routes))
}
