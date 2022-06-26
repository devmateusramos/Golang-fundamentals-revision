package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros) // veio um slice

	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

// Não pode ter mais de um variatico por função
// E ele tem que obrigatoriamente ser o último parâmetro

func escreverTexto(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
}

func main() {
	total := soma(2, 3, 4)
	fmt.Println(total)
	escreverTexto("Olha os números: ", 1, 2, 3, 4, 51)
}
