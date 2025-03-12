package handlers

import (
	"fmt"
	"net/http"
	"apirest/models"
	"encoding/json"
	"apirest/db"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	//rw.Header().Set("Content-Type", "text/xml")
	
	db.Connect()
	users := models.ListUsers()
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	
	rw.Header().Set("Content-Type", "application/json")
	
	// Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(userId)
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	
	rw.Header().Set("Content-Type", "application/json")
	
	// Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil{
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	}else{
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	
	rw.Header().Set("Content-Type", "application/json")
	
	// Actualizar registro
	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil{
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	}else{
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	
	rw.Header().Set("Content-Type", "application/json")
	
	// Elminar ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(userId)
	user.Delete()
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}