package main

import "fmt"

func main() {
	// declaración de constantes
	const pi float64 = 3.141592653589793
	const pi2 = 3.141592653589793
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// declaración de variabes enteras
	base := 12
	var altura int = 14
	var area int
	area = base * altura

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calculo area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("área cuadrado: ", areaCuadrado)
}
