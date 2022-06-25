package main

import "fmt"

func diaDaSemana(numero int) string {
	//switch {
	//case numero == 1:
	//	return "Domingo"
	//case numero == 2:
	//	return "Segunda"
	//case numero == 3:
	//	return "Terça"
	//default:
	//	return "Inválido"
	//}
	//switch numero {
	//case 1:
	//	return "Domingo"
	//case 2:
	//	return "Segunda"
	//case 3:
	//	return "Terça"
	//default:
	//	return "Inválido"
	//}

	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	default:
		diaDaSemana = "Inválido"
	}
	return diaDaSemana
}

// falltrough => ele vai avaliar a condiçao e se tiver nela falthrough
// Ele em vez de retornar a clausula dessa condiçao que ele entrou
// Ele vai retornar a clausula da proxima condiçao direto, sem avaliar ela

func main() {
	dia1 := diaDaSemana(3)
	fmt.Println(dia1)
}
