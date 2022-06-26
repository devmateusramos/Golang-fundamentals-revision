package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova() // Vai ser o texto de closure
	// Porque quando eu criei a variavel, texto referenciava a variavel texto
	// do escopo da closure
	// basicamente função closure é isso, de como se tivesse uma memória
	// de como ficou referenciado a função
}
