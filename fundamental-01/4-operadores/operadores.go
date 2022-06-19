package main

import "fmt"

func main() {
	// aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	// Não pode fazer operações com variáveis de tipos diferentes
	// Por exemplo, não pode somar int8 com int32
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	// operadores de atribuição
	var variavel string = "string"
	variavel2 := "string também"
	fmt.Println(variavel2, variavel)
	// operadores relacionais (sempre booleanos)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// operadores lógicos são 3
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	// operadores unários
	numeroParaAumentar := 10
	numeroParaAumentar++     // aumentou em 1
	numeroParaAumentar += 10 // ele mais 10
	numeroParaAumentar--     // diminui 1
	numeroParaAumentar -= 20 // ele menos 20
	// esses de ++ e -- só tem adição e subtração
	numeroParaAumentar *= 3  // Ele vezes 3
	numeroParaAumentar /= 10 // ele dividido por 10
	numeroParaAumentar %= 2  // Resto dele dividido por 2
	fmt.Println(numeroParaAumentar)
	// operador ternário não tem, somente tem o if e else normal
}
