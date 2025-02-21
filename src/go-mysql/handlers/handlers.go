package handlers

 import "database/sql"

 // ListContacts lista todos los contactos desde la base de datos
 func ListContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	// Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Iterar sobre los resultados y mostrarlos 
	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println()
 }