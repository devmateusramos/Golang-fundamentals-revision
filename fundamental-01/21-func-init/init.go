package main

import "fmt"

func init() {
	fmt.Println("Executando a função init")
}

// É uma função que roda antes da função main, a ordem dela no código
//não influencia,
//se colocar depois vai ser iniciada antes mesmo assim
// Você pode ter uma função init por arquivo (sim, arquivo e não pacote)
func main() {
	fmt.Println("Executando a função main")
}
