package custom_errors

import (
	"errors"
	"net/http"
)

type HttpError struct {
	Code        string
	Status      int
	Description string
}

func (he HttpError) Error() string {
	return he.Description
}

type ServiceError struct {
	Code        string
	Description string
}

func (se ServiceError) Error() string {
	return se.Description
}

//sentinel errors

var errorParsed = ServiceError{
	Code:        "InternalError",
	Description: "couldnt parse store response to go struct",
}
var errorNotfound = ServiceError{
	Code:        "NotFound",
	Description: "user not found",
}
var duplicatedIdError = ServiceError{
	Code:        "IdAlreadyInUse",
	Description: "id already used",
}
var conectionToStoreError = ServiceError{
	Code:        "ConectionError",
	Description: "connection refused",
}

func CreateHttpError(e error) HttpError {
	serviceError, ok := e.(ServiceError)
	if !ok {
		return HttpError{
			Code:        "InternalError",
			Status:      http.StatusInternalServerError,
			Description: e.Error(),
		}
	}
	switch {
	case errors.Is(e, conectionToStoreError):
		return HttpError{
			Code:        serviceError.Code,
			Status:      http.StatusInternalServerError,
			Description: serviceError.Description,
		}
	case errors.Is(e, errorParsed):
		return HttpError{
			Code:        serviceError.Code,
			Status:      http.StatusInternalServerError,
			Description: serviceError.Description,
		}
	case errors.Is(e, errorNotfound):
		return HttpError{
			Code:        serviceError.Code,
			Status:      http.StatusNotFound,
			Description: serviceError.Description,
		}
	case errors.Is(e, duplicatedIdError):
		return HttpError{
			Code:        serviceError.Code,
			Status:      http.StatusConflict,
			Description: serviceError.Description,
		}
	default:
		return HttpError{
			Code:        "InternalError",
			Status:      http.StatusInternalServerError,
			Description: "There was an unspecified internal error",
		}
	}

}

var Error_UserNotFound = errors.New("user not found")
var Error_UuidAlreadyExists = errors.New("there is already an user with this uui")

var Error_WrongBodyFormat = errors.New("wrong body format")
var Error_ParsingJson = errors.New("could not parse Json to User")
