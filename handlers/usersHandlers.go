package handlers

import (
	"bootcam1_users/custom_errors"
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
		sendError(w, http.StatusBadRequest, errParsing)
		return
	}

	user, err := usersManager.Get(parsedId)
	if err != nil {
		sendError(w, http.StatusNotFound, err)
		return
	}

	//parsear los usuarios a Json
	json.NewEncoder(w).Encode(user)
}
func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userRequest structures.UserRequest
	decoder := json.NewDecoder(r.Body)
	//check if body is in json format
	if err := decoder.Decode(&userRequest); err != nil {
		sendError(w, http.StatusBadRequest, custom_errors.Error_WrongBodyFormat)
		return
	}

	//create uuid
	newUuid := uuid.New()
	user := structures.User{
		ID:       newUuid,
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		Email:    userRequest.LastName,
		Active:   userRequest.Active,
		Address:  userRequest.Address,
	}

	createdUser, err := usersManager.Create(user)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
	//parsear los usuarios a Json
	json.NewEncoder(w).Encode(createdUser)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//insertar un header
	w.Header().Set("Content-Type", "application/json")

	parsedId, errParsing := uuid.Parse(id)
	if errParsing != nil {
		sendError(w, http.StatusBadRequest, errParsing)
		return
	}

	err := usersManager.Delete(parsedId)
	if err != nil {
		sendError(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	//parsear los usuarios a Json
	json.NewEncoder(w).Encode(map[string]string{
		"response": fmt.Sprintf("user with id %v deleted", id)})
}
func sendError(w http.ResponseWriter, status int, message error) {
	fmt.Printf("ERROR ON HANDLERS, %v", message)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf(message.Error()),
		"code":    fmt.Sprintf("%v", status)})
}
