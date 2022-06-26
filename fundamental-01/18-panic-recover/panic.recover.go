package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
	// ele vai recuperar a execução do programa
	// vai retornar false pq é o valo zero do retorno dessa função
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic(" A MÉDIA É EXATAMENTE 6")
	// diferente do erro, no erro o seu programa continua sendo executado
	// no panic ele mata seu programa se vc não tiver recover pra
	// recuperar ele
	// Antes de matar tudo e interroper a execução
	// Ele vai chamar todas os defer que tiverem
}

func main() {
	fmt.Println(alunoEstaAprovado(4, 8))
	fmt.Println("Pós execução")
}
