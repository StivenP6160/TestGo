package main

import (
	"fmt"
	"os"
)

//Ejemplo
/*
func main() {
	defer fmt.Println("Este mensaje se ejecutará al final")

	fmt.Println("Inicio del programa")
}
*/

func main() {
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Stiven...este es un archivo de prueba usando la palbra clave defer"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
