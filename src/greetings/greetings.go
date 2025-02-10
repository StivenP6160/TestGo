package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especifica
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("El nmombre está vacío")
	}
	//Devuelve un saludo que incluye el nombre en un mensaje.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

/*
funcion Hello cuyo parametro de entrada es name de tipo string (name string), que devuelve un string o un error (string error)
*/

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}
