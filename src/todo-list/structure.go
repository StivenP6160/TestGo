package main

import "fmt"

//Forma para crear una estructura
type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {

	var p Persona
	p.nombre = "Stiven"
	p.edad = 33
	p.correo = "stiven@example.com"

	fmt.Println(p)
	fmt.Println("**************************************")

	//Instanciar en una sola linea
	a := Persona{nombre: "Stiven", edad: 33, correo: "stiven@example.com"}

	fmt.Println(a.edad)
	fmt.Println("**************************************")

}
