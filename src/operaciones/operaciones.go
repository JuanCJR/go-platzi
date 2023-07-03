package operaciones

import "fmt"

//Operaciones is a function
func Operaciones() {
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
