package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(user)

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)

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
