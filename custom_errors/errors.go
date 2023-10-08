package custom_errors

import (
	"errors"
	"fmt"
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
	return fmt.Sprintf("service error: %q", se.Description)
}

// codes
const (
	Internal        = "InternalError"
	NotFound        = "NotFound"
	DuplicatedId    = "DuplicatedKey"
	ConectionFailed = "ConectionFailed"
)

//sentinel errors

var statusMap = map[string]int{
	Internal:        http.StatusInternalServerError,
	NotFound:        http.StatusNotFound,
	DuplicatedId:    http.StatusConflict,
	ConectionFailed: http.StatusInternalServerError,
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
	status, found := statusMap[serviceError.Code]
	if !found {
		return HttpError{
			Code:        "InternalError",
			Status:      http.StatusInternalServerError,
			Description: serviceError.Description,
		}
	}
	return HttpError{
		Code:        serviceError.Code,
		Status:      status,
		Description: serviceError.Description,
	}

}

var Error_UserNotFound = errors.New("user not found")
var Error_UuidAlreadyExists = errors.New("there is already an user with this uui")

var Error_WrongBodyFormat = errors.New("wrong body format")
var Error_ParsingJson = errors.New("could not parse Json to User")
