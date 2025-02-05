package main

import "fmt"

func main() {
	//Declarar un slice
	//var a []int

	//Inicializar un slice
	diaSemana := []string{"Domingo", "Lunes", "Martes",
		"Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println(diaSemana)

	//Crear un slice apartir de otro slice
	diaSlice := diaSemana[0:5]
	fmt.Println(diaSlice)

	//Agregar elementos
	diaSlice = append(diaSlice, "Viernes")
	fmt.Println(diaSlice)

	//Quitar elementos
	diaSlice = append(diaSlice[:2], diaSlice[3:]...)
	fmt.Println(diaSlice)

	//Longitud y Cantidad de elementos
	fmt.Println(len(diaSlice)) //Longitud
	fmt.Println(cap(diaSlice)) //Cantidad resultante es la de "diaSemana" ya que es un slice de "diaSemana"

}
