package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	// como já foi definido a variavel de retorno e não só o tipo
	// não pode inicializar ela e não precisa
	//apontar pro return que ela é o retorno
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
