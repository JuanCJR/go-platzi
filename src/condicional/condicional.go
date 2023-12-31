package condicional

import (
	"fmt"
	"log"
	"strconv"
)

// CondicionalFunc is a function
func CondicionalFunc() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With AND
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}

	//With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es Vedad, Or")
	}

	//Convertir Texto a numero
	value, err := strconv.Atoi("34")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value:", value)
}
