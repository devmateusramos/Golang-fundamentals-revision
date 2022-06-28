package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // Aqui vai ser -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()
	escrever("Fora das goroutines de função anônima")
	fmt.Println("Antes do waitGroup fechar")
	waitGroup.Wait() // 0
	fmt.Println("Ultima linha de código")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
