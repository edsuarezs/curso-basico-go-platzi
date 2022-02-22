package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	minus := strings.ToLower(text)
	var textReverse string
	for i := len(minus) - 1; i >= 0; i-- {
		textReverse += string(minus[i])
	}

	if minus == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}
	// este for imprime el indice del slice y el valor del indice del slice
	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	// este for imprime solo el valor del slice
	for _, valor := range slice {
		fmt.Println(valor)
	}
	// este for imprime solo el indice del slice
	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("Ama")
}
