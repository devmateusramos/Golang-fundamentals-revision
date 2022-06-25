package main

import (
	"fmt"
)

func main() {
	//i := 0
	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)
	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementando J")
	//	time.Sleep(time.Second)
	//	fmt.Println(j)
	//}
	//// Agora j não existe fora desse for, ela só existe no escopo dele
	//for j := 0; j < 10; j += 3 {
	//	fmt.Println("Incrementando J")
	//	time.Sleep(time.Second)
	//	fmt.Println(j)
	//}
	//
	//palavras := [3]string{"Mateus", "Engenheiro de", "Software"}
	//for indice, valor := range palavras {
	//	fmt.Println("indice: ", indice)
	//	fmt.Println("valor: ", valor)
	//}
	// funciona igual pra slices
	for indice, letra := range "PALAVRA SENDO ITERADA" {
		fmt.Println(indice, string(letra))
	}

	usuarioMap := map[int]int{
		1: 2,
		2: 4,
		3: 5,
	}
	for chave, valor := range usuarioMap {
		fmt.Println(chave, valor)
	}

	// não dá pra fazer range em struct

	for {
		fmt.Println("Assim faz um loop infinito com for")
	}
}
