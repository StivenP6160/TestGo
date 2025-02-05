package main

import (
	"fmt"
)

//Declaración de constantes

const Pi = 3.14

//Declaración de varios constantes a la vez

const (
	domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	/*
			fmt.Println("Hola mundo")
			fmt.Println(quote.Go())

			fullName := "Alex Roel \t(alias \"roelcode\")\n"
			fmt.Println(fullName)


		var a byte = 'a'
		fmt.Println(a)

		s := "hola"
		fmt.Println(s[0])

		//Declaración de variables


			var firstName, lastName string
			var age int

			firstName = "Stiven"
			lastName = "Patiño"
			age = 33
	*/

	//Otrta manera de inicializar variables
	/*
					firstName, lastName ,age := "Stiven", "Patiño", 33

					fmt.Println(firstName, lastName , age)

					fmt.Println(Pi)

					fmt.Println(Viernes)


				var (
					defaultInt    int
					defaultUint   uint
					defaultFloat  float32
					defaultBool   bool
					defaultString string
				)

				fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)


			var integer16 int16 = 50
			var integer32 int32 = 150


		z := "100"
		i, _ := strconv.Atoi(z)

		fmt.Println(i)
	*/

	name := "Stiven"
	age := 33

	fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age)
}
