package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Iniciando para dizer se passou ou não")
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}

	return false
}

func main() {
	// Ela adia a execução de algo que você criou
	// Significa adiar, então ele vai adiar a execução
	// até o último momento possível
	defer funcao1()
	funcao2()
	// Se tivesse retorno seria executando imediatamente antes do retorno
	// ser dado
	fmt.Println(alunoEstaAprovado(5, 8))
}
