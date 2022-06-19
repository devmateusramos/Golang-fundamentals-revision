package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// tem o tipo func() que você pode atribuir para uma variável, armazenando nela uma função
// funções em go podem ter mais de um retorno
// para ignorar um do retorno, você pode colocar o _
// a ordem importa!!
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(23, 21)
	fmt.Println(soma)
	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("Ola mundo pela função  na variavel f tipo func()")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	resultaSomaUnico, _ := calculosMatematicos(12, 31)
	fmt.Println(resultaSomaUnico)
}
