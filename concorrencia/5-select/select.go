package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		//mensagemCanal1 := <-canal1 // Aqui ele vai podia escrever 4 canal 1
		//// Antes do canal 2, mas perde tempo desnecessário esperando o código
		//// de baixo acontecer
		//// pra isso usamos o select
		//fmt.Println(mensagemCanal1)
		//mensagemCanal2 := <-canal2
		//fmt.Println(mensagemCanal2)
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
