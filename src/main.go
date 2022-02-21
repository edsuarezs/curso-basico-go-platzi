package main

import "fmt"

func main() {

	//Declaración de variables
	helloMessage := "Hello "
	worldMessage := "World"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) //Mientras lo sepa una buena practica es agregar el tipo de datos

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("cursos: %T \n", cursos)
}
