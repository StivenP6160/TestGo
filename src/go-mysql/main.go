package main

import (
	"fmt"
	"log"
	"go-mysql/database"
	_ "github.com/go-sql-driver/mysql"
	"go-mysql/handlers"
	//"go-mysql/models"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()


	/*
	// Crear una instancia de Contact
	newContact := models.Contact{
		Id: 3,
		Name: "Carlo López",
		Email: "carlosLopez@example.com",
		Phone: "321-654-0987",
	}
	
	handlers.DeleteContact(db, 4)

	handlers.ListContacts(db)
	*/
	
	// Que sea interactivo
	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contactos por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Println("Seleccione la opción: ")

		// Leer la opción seleccionada por el usuario
		var option int
		fmt.Scanln(&option)

	}
}