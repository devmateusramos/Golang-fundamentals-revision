package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Mateus",
		"sobrenome": "Ramos",
	}
	fmt.Println(usuario)
	// para acessar tem que usar o colchete
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]int8{
		"nome": {
			"primeiroNumero": 2,
		},
	}

	fmt.Println(usuario2)
	delete(usuario, "nome")
	fmt.Println(usuario)
}
