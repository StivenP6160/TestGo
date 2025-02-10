package greetings

import "fmt"

//Hello devuelve un saludo para la persona especificada.
func Hello(name string) string {
	//Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf("¡Hola, %v! ¡Bienvenido!", name)
	return message
}
