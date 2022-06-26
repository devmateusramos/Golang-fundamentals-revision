package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(1)
	generica("string sendo usada")
	generica(true)
	// cuidado no uso, dependendo perde o sentido como no exemplo abaixo
	mapa := map[interface{}]interface{}{
		1:        "string",
		"string": true,
	}
	fmt.Println(mapa)
}
