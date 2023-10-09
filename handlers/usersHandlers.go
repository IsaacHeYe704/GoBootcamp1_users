package handlers

import (
	"bootcam1_users/custom_errors"
	"bootcam1_users/service"
	"bootcam1_users/structures"
	"encoding/json"
	"errors"
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
			sendHttpError(w, castHttpError(err))
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
			sendHttpError(w, custom_errors.HttpError{
				Code:        custom_errors.WrongBodyFormat,
				Status:      http.StatusBadRequest,
				Description: errParsing.Error(),
			})
			return
		}

		user, err := service.Get(parsedId)
		if err != nil {
			sendHttpError(w, castHttpError(err))
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
			sendHttpError(w, custom_errors.HttpError{
				Code:        custom_errors.WrongBodyFormat,
				Status:      http.StatusBadRequest,
				Description: err.Error(),
			})
			return
		}
		//validate the user received
		validate := validator.New()
		structError := validate.Struct(userRequest)
		if structError != nil {
			sendHttpError(w, custom_errors.HttpError{
				Code:        custom_errors.WrongBodyFormat,
				Status:      http.StatusBadRequest,
				Description: structError.Error(),
			})
			return
		}

		//validate if the user could be stored
		createdUser, err := service.Create(userRequest)
		if err != nil {
			//convert service error into http error
			httpError := castHttpError(err)
			//send error
			sendHttpError(w, httpError)
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
			sendHttpError(w, custom_errors.HttpError{
				Code:        custom_errors.WrongBodyFormat,
				Status:      http.StatusBadRequest,
				Description: errParsing.Error(),
			})
			return
		}

		var userRequest structures.UserRequest
		decoder := json.NewDecoder(r.Body)
		//check if body is in json format
		if err := decoder.Decode(&userRequest); err != nil {
			sendHttpError(w, custom_errors.HttpError{
				Code:        custom_errors.WrongBodyFormat,
				Status:      http.StatusBadRequest,
				Description: err.Error(),
			})
			return
		}

		//validate the user received
		validate := validator.New()
		structError := validate.Struct(userRequest)
		if structError != nil {
			sendHttpError(w, castHttpError(structError))
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
			sendHttpError(w, castHttpError(err))
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
			sendHttpError(w, custom_errors.HttpError{
				Code:        custom_errors.WrongBodyFormat,
				Status:      http.StatusBadRequest,
				Description: errParsing.Error(),
			})
			return
		}
		validator.New()

		err := service.Delete(parsedId)
		if err != nil {
			sendHttpError(w, castHttpError(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		//parsear los usuarios a Json
		json.NewEncoder(w).Encode(map[string]string{
			"response": fmt.Sprintf("user with id %v deleted", id)})
	}
	return fn
}

func castHttpError(err error) custom_errors.HttpError {
	errAssert, ok := err.(custom_errors.ServiceError)

	if !ok {
		return custom_errors.HttpError{
			Code:        errAssert.Code,
			Status:      http.StatusInternalServerError,
			Description: errAssert.Error(),
		}
	}

	if errors.Is(errAssert, custom_errors.ServiceError{Code: custom_errors.NotFound}) {
		return custom_errors.HttpError{
			Code:        errAssert.Code,
			Status:      http.StatusNotFound,
			Description: errAssert.Error(),
		}
	}
	if errors.Is(errAssert, custom_errors.ServiceError{Code: custom_errors.ConectionFailed}) {
		return custom_errors.HttpError{
			Code:        errAssert.Code,
			Status:      http.StatusInternalServerError,
			Description: errAssert.Error(),
		}
	}
	if errors.Is(errAssert, custom_errors.ServiceError{Code: custom_errors.DuplicatedId}) {
		return custom_errors.HttpError{
			Code:        errAssert.Code,
			Status:      http.StatusConflict,
			Description: errAssert.Error(),
		}
	}
	return custom_errors.HttpError{
		Code:        errAssert.Code,
		Status:      http.StatusInternalServerError,
		Description: errAssert.Error(),
	}
}
func sendHttpError(w http.ResponseWriter, httpError custom_errors.HttpError) {
	slog.Error("returning error ", "statusCode", fmt.Sprint(httpError.Status), "message", fmt.Sprint(httpError.Error()))
	w.WriteHeader(httpError.Status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": httpError.Error(),
		"code":    fmt.Sprintf("%v", httpError.Code)})
}
