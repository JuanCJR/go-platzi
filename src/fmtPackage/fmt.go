package fmtPackage

import (
	"fmt"
)

// FmtFunc is a function
func FmtFunc() {
	// Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)

	//%v se utiliza cuando no se conoce el tipo de dato
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf guarda el mensaje en una variable como string
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Imprimir tipo de datoTipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
