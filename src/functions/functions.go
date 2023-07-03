package functions

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {

	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

//FunctionsFunc is a function
func FunctionsFunc() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println(value)
	value1, value2 := dobleReturn((2))

	fmt.Println(value1, value2)

}
