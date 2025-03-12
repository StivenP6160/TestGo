package main

import (
	"gomysql/db"
	"gomysql/models"
	"fmt"
)

func main() {
	db.Connect()
	
	//fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")

	// Código para crear el registro
	// user := models.CreateUser("Stiven", "patino6160", "stiven@example.com")
	// fmt.Println(user)

	//Código para listar todos los registros
	users := models.ListUsers()
	fmt.Println(users)

	//Código para listar un usuario especifico
	//user := models.GetUser(2)
	//fmt.Println(user)

	//Código para actualizar un registro
	//user.Username = "Juan"
	//user.Password = "juan0987"
	//user.Email = "juan@example.com"
	//user.Save()

	//Eliminar un registro (Se borrará el registro que estoy visualizando en GetUser)
	//user.Delete()

	//TruncateTable borra todos los registros y reinicia la tabla 
	//db.TruncateTable("users")

	fmt.Println(models.ListUsers())

	db.Close()
	//db.Ping()
}