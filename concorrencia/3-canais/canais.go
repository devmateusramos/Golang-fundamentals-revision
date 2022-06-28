package main

import (
	"fmt"
	"time"
)

func main() {
	//Canal de comunicação, pode enviar ou receber dados,
	// usado para sincronizar as goroutines
	canal := make(chan string)
	go escrever("Ola mundo", canal)
	fmt.Println("Depois da função escrever começar a ser executada")
	mensagem := <-canal // Nesse caso vai ocorrer apenas uma vez e vai acabar
	// a função, pois ele vai dar prosseguimento com a mensagem
	// assim que tiver um valor em canal
	// Se colocar essa mensagem num for infinito
	// o canal vai ficar esperando as 5x, mas dps vai continuar e nunca algo
	// vai ser passado, ocasionando num deadlock
	fmt.Println(mensagem)
	// Tem que ter cuidado com deadlock porque ele não é pego em compilação
	// Só em execução
	// A maneira de tratar isso é essa abaixo, lidando com se ele está
	// aberto ou fechado, aberto ainda vai enviar/receber dados, fechado não
	//for {
	//	mensagem2, aberto := <-canal
	//	if !aberto {
	//		break // uma maneira de sair do loop infinito em for é chamando o
	//		// break
	//	}
	//	fmt.Println(mensagem2)
	//}
	// Maneira mais legal de fazer isso
	for mensagem2 := range canal { // Assim que identificar q canal fechou
		// Ele para de esperar valores
		fmt.Println(mensagem2)
	}
	fmt.Println("Fim do programa")
}
func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) // aqui marcamos que vai ser fechado o canal, assim é
	// possível receber isso em mensagem
}
