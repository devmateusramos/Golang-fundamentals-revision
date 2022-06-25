package main

import "fmt"

func main() {
	numero := 10
	if numero < 15 {
		fmt.Println("É menor que 15")
	} else {
		fmt.Println("É maior que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	}
	// Essa variável não esta definida fora do escopo do if,
	//por isso não é possível de ser acessada fora dele, vai apontar o erro
	// de variável não definida
	if outroNumero := numero; outroNumero < 0 {
		fmt.Println("Número é menor que zero")
	} else {
		fmt.Println("Número é maior que zero")
	}
}
