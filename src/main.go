package main

import (
	"fmt"
)

func variables() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}

func operaciones() {
	// Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// suma
	result := x + y
	fmt.Println("Suma:", result)

	// resta
	result = y - x
	fmt.Println("Resta:", result)

	// multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// division
	result = y / x
	fmt.Println("Division:", result)

	// modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// incremental
	x++
	fmt.Println("Incremental:", x)

	// decremental
	x--
	fmt.Println("Decremental:", x)
	// rectangulo
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area rectangulo:", areaRectangulo)
}

func main() {
	variables()
	operaciones()
	FmtTest()
}
