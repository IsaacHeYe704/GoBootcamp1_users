package handlers

import (
	"bootcam1_users/custom_errors"
	"bootcam1_users/service"
	"bootcam1_users/structures"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllUsers(service service.UserService) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		users, err := service.GetAll()
		//insertar un header
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			sendError(w, custom_errors.CreateHttpError(err))
			return
		}
		//preguntar a Ever
		w.WriteHeader(http.StatusFound)
		//parsear los usuarios a Json
		json.NewEncoder(w).Encode(users)
	}

	return fn
}

func GetUserById(service service.UserService) func(w http.ResponseWriter, r *http.Request) {

	fn := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		//insertar un header
		w.Header().Set("Content-Type", "application/json")

		parsedId, errParsing := uuid.Parse(id)
		if errParsing != nil {
			sendError(w, custom_errors.CreateHttpError(errParsing))
			return
		}

		user, err := service.Get(parsedId)
		if err != nil {
			sendError(w, custom_errors.CreateHttpError(err))
			return
		}

		//parsear los usuarios a Json
		json.NewEncoder(w).Encode(user)
	}
	return fn
}
func PostUser(service service.UserService) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//json marshall
		var userRequest structures.UserRequest
		decoder := json.NewDecoder(r.Body)
		//check if body is in json format
		if err := decoder.Decode(&userRequest); err != nil {
			httpError := custom_errors.CreateHttpError(err)
			sendError(w, httpError)
			return
		}
		//validate the user received
		validate := validator.New()
		structError := validate.Struct(userRequest)
		if structError != nil {
			httpError := custom_errors.CreateHttpError(structError)
			sendError(w, httpError)
			return
		}

		//validate if the user could be stored
		createdUser, err := service.Create(userRequest)
		if err != nil {
			//convert service error into http error
			httpError := custom_errors.CreateHttpError(err)
			//send error
			sendError(w, httpError)
			return
		}

		//parsear los usuarios a Json
		json.NewEncoder(w).Encode(createdUser)
	}
	return fn
}
func PutUsers(service service.UserService) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		//check if the id is a valid uuid
		parsedId, errParsing := uuid.Parse(id)
		if errParsing != nil {
			sendError(w, custom_errors.CreateHttpError(errParsing))
			return
		}

		var userRequest structures.UserRequest
		decoder := json.NewDecoder(r.Body)
		//check if body is in json format
		if err := decoder.Decode(&userRequest); err != nil {
			sendError(w, custom_errors.CreateHttpError(err))
			return
		}

		//validate the user received
		validate := validator.New()
		structError := validate.Struct(userRequest)
		if structError != nil {
			sendError(w, custom_errors.CreateHttpError(structError))
			return
		}

		//parsed structs from userRest to user
		user := structures.User{
			ID:       parsedId,
			Name:     userRequest.Name,
			LastName: userRequest.LastName,
			Email:    userRequest.LastName,
			Active:   userRequest.Active,
			Address:  userRequest.Address,
		}
		//validate if the user could be stored
		updatedUser, err := service.Update(parsedId, user)
		if err != nil {
			sendError(w, custom_errors.CreateHttpError(err))
			return
		}
		json.NewEncoder(w).Encode(updatedUser)
	}
	return fn
}
func DeleteUsers(service service.UserService) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		//insertar un header
		w.Header().Set("Content-Type", "application/json")

		parsedId, errParsing := uuid.Parse(id)
		if errParsing != nil {
			sendError(w, custom_errors.CreateHttpError(errParsing))
			return
		}
		validator.New()

		err := service.Delete(parsedId)
		if err != nil {
			sendError(w, custom_errors.CreateHttpError(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		//parsear los usuarios a Json
		json.NewEncoder(w).Encode(map[string]string{
			"response": fmt.Sprintf("user with id %v deleted", id)})
	}
	return fn
}
func sendError(w http.ResponseWriter, httpError custom_errors.HttpError) {
	slog.Error("returning error ", "statusCode", fmt.Sprint(httpError.Status), "message", fmt.Sprint(httpError.Error()))
	w.WriteHeader(httpError.Status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf(httpError.Error()),
		"code":    fmt.Sprintf("%v", httpError.Code)})
}
