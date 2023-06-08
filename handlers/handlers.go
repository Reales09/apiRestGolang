package handlers

import (
	"apirest/db"
	"apirest/models"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	// rw.Header().Set("Content-Type", "application/json")
	// rw.Header().Set("Content-Type", "test/xml")
	db.Connect()
	users := models.ListUsers()
	db.Close()
	output, _ := yaml.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtiene un usuario")
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Crea un usuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Actualiza un usuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Elimina un usuario")
}
