package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := banco.Conection()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewRepositoryUsers(db)
	repository.Create(user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando todos os usuários")
}

func SearchUser(w http.ResponseWriter, r *http.Request) {

}

func ChangeUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
