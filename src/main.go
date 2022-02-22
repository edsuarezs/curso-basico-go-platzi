package main

import "fmt"

func main() {
	// Defer, es el keyword para ejecutar la última función antes de que todo muera.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
