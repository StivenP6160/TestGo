package main

import "fmt"

func main() {
	hello("Stiven")

	fmt.Println("******************************")

	//Para la función que retorna
	retorno := saludo("John")
	fmt.Println(retorno)

	fmt.Println("******************************")

	//Función de multiples valores
	sum, mul := calc(4, 5)
	fmt.Println("La suma es:", sum)
	fmt.Println("La multiplicación es:", mul)

}

func hello(name string) {
	fmt.Println("Hola,", name)

}

//Función que retorna valores
func saludo(name string) string {
	return "Hola, " + name
}

//Función de multiples valores
func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}
