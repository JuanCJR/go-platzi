package switchPackage

import "fmt"

//SwitchFunc is a function
func SwitchFunc() {

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

	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}
