// Ejemplo

package main

import (
	"fmt"
	"math"
)

func main() {

	var num1, num2 float64

	fmt.Print("Ingrese el Cateto A: ")
	fmt.Scan(&num1)

	fmt.Print("Ingrese el Cateto B: ")
	fmt.Scan(&num2)

	hipotenusa := math.Sqrt(math.Pow(float64(num1), 2) + math.Pow(float64(num2), 2))
	fmt.Printf("El valor de la hipotenusa es: %.2f", hipotenusa)

	area := (num1 * num2) / 2
	fmt.Printf("El area del triangulo es: %.2f", area)

	perimetro := num1 + num2 + float64(hipotenusa)
	fmt.Printf("El perimetro es: %.2f", perimetro)

	/*
		var num1 int
		fmt.Print("Ingrese el primer número: ")
		fmt.Scan(&num1)
		potenciaNum1 := math.Pow(float64(num1), 2)
		fmt.Printf("%.1f", potenciaNum1)
	*/
}
