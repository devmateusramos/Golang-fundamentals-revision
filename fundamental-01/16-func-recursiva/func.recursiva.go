package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	// 1 1 2 3 5 8 13
	// Não é muito utilizada pra casos muito grandes, pra não dar um stackoverflow
	valorSequenciaFibonacciNaPosicao := fibonacci(4)
	fmt.Println(valorSequenciaFibonacciNaPosicao)
	for i := uint(1); i <= 20; i++ {
		fmt.Println(fibonacci(i))
	}
}
