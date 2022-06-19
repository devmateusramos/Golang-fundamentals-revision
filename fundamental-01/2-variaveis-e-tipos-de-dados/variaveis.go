package main

import (
	"errors"
	"fmt"
)

func main() {
	// Existem 4 tipos de números inteiros no Go
	// int8, int16, int32, int64
	var numeroInt8 int8 = 12
	fmt.Println(numeroInt8)
	var numeroInt16 int16 = 12
	fmt.Println(numeroInt16)
	var numeroInt32 int32 = 12
	fmt.Println(numeroInt32)
	var numeroInt64 int64 = 12
	fmt.Println(numeroInt64)
	// pode passar int sozinho também, nesse caso ele vai usar a arquitetura do seu computador como base
	var numeroIntSozinho int = 12
	fmt.Println(numeroIntSozinho)
	// Se você usar inferência de tipos ele vai adotar o tipo int que é esse da arquitetura do computador
	// int32, se tipar com rune então está tipando com int32, é apenas um alias
	numero := 10000
	fmt.Println(numero)
	var numeroRune rune = 13
	fmt.Println(numeroRune)
	// uint => unsygned int tem os mesmos que o int mas sem sinal tem um alias chamado rune que é o mesmo que o
	var numeroUint8 uint8 = 12
	fmt.Println(numeroUint8)
	// uint 8 tem o alias que você pode colocar que é o byte
	var numeroByte byte = 12
	fmt.Println(numeroByte)
	// falando de números reais tem o float32 e o float64
	var numeroFloat64 float64 = 12.45
	fmt.Println(numeroFloat64)
	// você não pode declarar só float, mas se usar inferência de tipo ele vai colocar somente float
	numeroFloatInfere := 12.54
	fmt.Println(numeroFloatInfere)
	// e nesse caso é o mesmo esquema da arquitetura do seu computador, se for 64 vai ser float64

	// string sempre aspas duplas, aspas simples ele considera da tabela asc
	var str string = "Texto"
	fmt.Println(str)
	// usando aspas simples ele retorna a representação da tabela asc
	tipoChar := 'o'
	fmt.Println(tipoChar)
	// valor zero => É o valor que vai ser atribuido para uma variável quando você não inicializa ela
	// para strings é string vazia, para números é o 0, para erro é nil.
	// No go todo tipo de dado tem seu valor inicial
	var texto string
	fmt.Println(texto) // Vai retornar valor em branco
	var numeroInt int16
	fmt.Println(numeroInt) // Vai retornar o valor zero
	// booleano tipo bool, true ou false, sendo que o valor zero(inicial) é o false
	var booleano bool = true
	fmt.Println(booleano)
	// tipo error, valor zero(inicial) é o nil, para criar o erro usamos o pacote errors
	var erro error
	fmt.Println(erro)
	var erroInicializado = errors.New("Erro interno")
	fmt.Println(erroInicializado)

}
