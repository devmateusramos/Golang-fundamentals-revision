package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	// tarefas concorrentes não necessariamente executam ao msm tempo
	// Se vc tiver um processador com só 1 núcleo, também poderia aplicar
	// concorrência
	go escrever("Olá mundo") // goroutine
	escrever("Programando em Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
