package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1 // Tem que usar o *, pois nesse caso se trata
	// do valor, sem ele vai estar mexendo com a referência, e não com o valor
	// que se encontra na referência
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero) // continua 20, só está sendo passado uma cópia
	numero2 := 20
	fmt.Println(numero2)
	inverterSinalComPonteiro(&numero2)
	fmt.Println(numero2)
}
