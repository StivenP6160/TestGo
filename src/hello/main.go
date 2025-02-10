package main

import (
	"fmt"
	"log"

	"github.com/StivenP6160/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Stiven", "Ana", "Andres"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	//message, err := greetings.Hello("Stiven")
	/*
		Se llama a la funcion Hello del paquete greetings
		Esta función devuelve dos valores: message y err
		El argumento es el nombre de la persona
	*/

	//if err != nil {
	//log.Fatal(err)
	/*
		Manejo de errores
		Esto verifica si greetings.Hello devuelve un error
		Si hay un error se usa log.Fatal() pra imprimir el error y detener la ejecución
		log.Fatal() es similar a log.Print, pero además de imprimir el mensaje, termina el programa.
	*/
	//}

	fmt.Println(messages)
	//Si no hubo error, se imprime el mensaje devuelto por la función greetings,Hello
}

/*
Si name es una cadena vacía, devuelve un error.
Si name no está vacío, devuelve un mensaje de saludo
*/
