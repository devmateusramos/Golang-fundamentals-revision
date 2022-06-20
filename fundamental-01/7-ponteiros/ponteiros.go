package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)
	// Por padrão as atribuições fazem uma cópia
	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3)
	fmt.Println(ponteiro)
	// Valor zero do ponteiro é nil
	// não pode colocar variavel 3 em ponteiro, pois ela não vai ser do tipo
	// ponteiro de int
	variavel3 = 100
	ponteiro = &variavel3 // para armazenar deve usar o &
	// Nesse caso está passando o endereço de memória da variavel3
	fmt.Println(variavel3)
	fmt.Println(ponteiro) // vai mostrar 0xc0000180e0
	// pois mostra o endereço na memória
	// para ver o valor dele utiliza o *
	fmt.Println(variavel3)
	fmt.Println(*ponteiro) // desreferenciação
}
