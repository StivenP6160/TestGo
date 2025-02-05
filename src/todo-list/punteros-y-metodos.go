package main

import "fmt"

//Punteros
/*
func main() {

		var x int = 10
		var p *int = &x

		fmt.Println(&x)
		fmt.Println(p)


	//Recibir un puntero
	var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)

}
//Desde otra funcion se modifica el valor de x
func editar(x *int) {
	*x = 20
}
*/

//Metodos
type Persona struct {
	nombre string
	edad   int
	ciudad string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es:", p.nombre)
}

func main() {
	p := Persona{"Stiven", 33, "stiven@example"}
	p.diHola()
}
