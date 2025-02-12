package animal

import "fmt"

type Animal interface {
	Sonido() //metodo
}

// Estructura de perro y sus métodos
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace gua guau")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau miua")
}

// Función para imprimir sonido
func HacerSonido(animal Animal) {
	animal.Sonido()
}