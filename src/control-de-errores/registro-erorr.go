package main

import (
	"log"
	"os"
)

/*
func main() {
	log.SetPrefix("registro():")
	log.Print("¡Oye, soy un log!")
	log.Fatal("¡Oye, soy un registro de erores!")
}
*/

// Registro de errores en una archivo
func main() {
	log.SetPrefix("main():")
	file, err := os.OpenFile("/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Oye, soy un log!")
}
