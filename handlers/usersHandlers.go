package handlers

import (
	"bootcam1_users/structures"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var usersManager = structures.NewUserManager()

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := usersManager.GetAll()
	//insertar un header
	w.Header().Set("Content-Type", "application/json")
	//preguntar a Ever
	w.WriteHeader(http.StatusFound)
	//parsear los usuarios a Json
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//insertar un header
	w.Header().Set("Content-Type", "application/json")

	parsedId, errParsing := uuid.Parse(id)
	if errParsing != nil {
		fmt.Printf("error %v", errParsing)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf(errParsing.Error())})
		return
	}

	user, err := usersManager.Get(parsedId)
	if err != nil {
		fmt.Printf("error %v", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf(err.Error())})
		return
	}

	//parsear los usuarios a Json
	json.NewEncoder(w).Encode(user)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//insertar un header
	w.Header().Set("Content-Type", "application/json")

	parsedId, errParsing := uuid.Parse(id)
	if errParsing != nil {
		fmt.Printf("error %v", errParsing)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf(errParsing.Error())})
		return
	}

	err := usersManager.Delete(parsedId)
	if err != nil {
		fmt.Printf("error %v", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf(err.Error())})
		return
	}

	w.WriteHeader(http.StatusOK)
	//parsear los usuarios a Json
	json.NewEncoder(w).Encode(map[string]string{
		"response": fmt.Sprintf("user with id %v deleted", id)})
}
