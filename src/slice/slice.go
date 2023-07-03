package slice

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

// SliceFunc is a function
func SliceFunc() {

	slice := []string{"hola", "que", "hace"}
	for _, valor := range slice {
		fmt.Println(valor)
	}
	isPalindromo("Ama")

}
