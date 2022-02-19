package main

import "fmt"

func main() {

	// Calculo area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("área cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = x - y
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// Division
	result = x / y
	fmt.Println("Division: ", result)

	// Modulo
	result = x % y
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental", x)

	// Reto: Cálcular el area del rectángulo, trapecio y de un circulo

	const pi float64 = 3.14159
	var radio float64
	var circunferencia float64 = 19
	base := x
	altura := y
	// área rectángulo
	result = base * altura
	fmt.Println("El área del rectángulo es: ", result)
	// área del trapecio
	base2 := base + x
	result = ((base2 + base) * altura) / 2
	fmt.Println("El área del trapecio es: ", result)
	// área del circulo
	radio = circunferencia / 2
	result2 := float64(result)
	result2 = pi * (radio * radio)
	fmt.Println("El área del circulo es: ", result2)

}
