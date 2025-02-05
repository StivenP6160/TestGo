package main

import "fmt"

func main() {
	/*
			var a [5]int
			a[0] = 10
			a[1] = 20

		//Inicializar matriz
		var a = [5]int{10, 20, 30, 40, 50}
	*/

	//Si no sé con cuantos valores va la matriz
	var a = [...]int{10, 20, 30, 40, 50}
	fmt.Println(a)

	//acceder a un elemento
	//fmt.Println(a[1])

	//iterar matriz
	/*
		for i := 0; i < len(a); i++ {
			fmt.Println(a[i])
		}
	*/

	for index, value := range a {
		fmt.Printf("Inidice: %d, Valor: %d\n", index, value)
	}

	//Matriz Bidimensional (Cada llave es una fila y cada objeto una columna)
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz)

}
