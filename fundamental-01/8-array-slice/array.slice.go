package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]string
	fmt.Println(array1) // como não coloquei nada, vai ter o valor 0 (string vazia)
	// em 5 itens
	array1[0] = "Posição 1"
	fmt.Println(array1)
	// Sempre tem que passar o tamanho para ser um array
	array2 := [3]string{"Posição 1", "Posição 2", "Posição3"}
	fmt.Println(array2)
	// Arrays tem tamanho fixo, não é possível adicionar além do já definido
	// O mais flexível possível é passar as reticências
	// E ele ir populando pelo que você passa
	array3 := [...]string{"posição1", "posição2", "posição3", "posição4"}
	fmt.Println(array3)

	// O Slice é uma estrutura baseada no array, mas com flexibilidades a mais
	// Principal diferença para criação é que vc não especifica nada,
	//só passa os colchetes
	slice := []int{12, 13, 14}
	fmt.Println(slice)
	// ele aponta pra um array, ele tem um componente dele que referencia isso
	fmt.Println(reflect.TypeOf(slice))  // Não tem limitação do tamanho
	fmt.Println(reflect.TypeOf(array3)) // Tem definido o tamanho
	// mostram tipos diferentes nisso
	slice = append(slice, 18)
	fmt.Println(slice)    // [12 13 14 18]
	slice2 := array3[1:2] // o 1 é inclusivo e o 3 é exclusivo, vai parar antes dele
	fmt.Println(slice2)   // trouxe apenas o [posição2]
	// mas isso é como se fosse um ponteiro apontando pra esses
	// índices do array
	array3[1] = "posição alterada na 1 do array 3"
	fmt.Println(slice2)

	// Se ele é um ponteiro e referenciando arrays
	// como ele faz quando vc cria ele sem ser a partir de um array?
	//Com arrays internos!
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
	slice3 = append(slice3, 4)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade Agora a capacidade foi para 24
	// Pois ele cria outro array para referência as novas informaçoes quando estoura
	// a capacidade anterior
	// Ele nunca vai deixar com que você fique sem espaço
	// no make o último argumento que e capacidade e opcional
	slice4 := make([]float64, 13)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	// Ele cria com a capacidade de ate onde você estabeleceu
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4)) // Aqui já aumentou a capacidade para 26
}
