package handlers

 import (
	"database/sql"
	"log"
	"fmt"
	"go-mysql/models"
)

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
	fmt.Println("-----------------------------------------------------------------------------------")
	for rows.Next() {
		// Instancia del modelo contact
		contact := models.Contact{}

		var valueEmail sql.NullString 
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		}else {
			contact.Email = "Sin correo electrónico"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("-----------------------------------------------------------------------------------")
	}
 }

 // GetContactByID obtiene un contacto de la base de datos mediante si ID
 func GetContactByID(db *sql.DB, contactID int){
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	// Instancia de modelo de contact
	contact := models.Contact{}
	var valueEmail sql.NullString // Para manejar valor null. 

	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows{
			log.Fatalf("No se encotró ningún contacto con el ID %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	}else {
		contact.Email = "Sin correo electrónico"
	}

	fmt.Println("\nLISTA DE CONTACTO:")
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
	contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("-----------------------------------------------------------------------------------")

}

// createContact registra un nuevo contacto en la base de datos
func CreateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQl para insertar un nuevo contacto
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}

// UpdateContact actualiza un contacto existente en la base de datos
func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto actualizado con éxito")
}

func DeleteContact(db *sql.DB, contactID int) {
	// Sentencia SQL para eliminar un contacto por su ID
	query := "DELETE FROM contact WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto eliminado con éxito")
}