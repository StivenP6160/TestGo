package main

import (
    //"library/book"
    "library/animal"
)

func main() {
    //myBook := book.NewBook("Moby Dick", "Herman Melville", 300)

    //myTextBook := book.NewTextBook("Comunicación", "Jaime Gamarra", 261, "Santillana SAC", "Secundaria")

    //myBook.PrintInfo()
    //myTextBook.PrintInfo()

    //book.Print(myBook)
    //book.Print(myTextBook)

    animales := []animal.Animal {
        &animal.Perro{Nombre:"Max"},
        &animal.Gato{Nombre:"Tom"},        
        &animal.Perro{Nombre:"Buddy"},
        &animal.Gato{Nombre:"Luna"},
    }
    // animales: es un slice que almacena objetos de tipo animal.Animal
    // animal.Animal: es una interfaz que define un conjunt de métodos que deben ser implementados por cualquier tipo que quiera ser considerado un "animal"

    for _, animal := range animales {
        animal.Sonido()
    }
}

    