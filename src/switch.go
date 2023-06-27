package main

import "fmt"

func main() {

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es Impar")
	}

	//switch sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 100:
		fmt.Println("Es menor a 100")
	default:
		fmt.Println("No condiciono")
	}

}
