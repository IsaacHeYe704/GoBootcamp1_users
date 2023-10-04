package custom_errors

import (
	"net/http"
)

// error wraping
var HttpError_notFound = HttpError{
	Code:        "NotFound",
	Status:      http.StatusNotFound,
	Description: "User not found",
}

var HttpError_idAlreadyExists = HttpError{
	Code:        "idAlreadyExists",
	Status:      http.StatusConflict,
	Description: "Conflict, id already in use.",
}

var HttpError_WrongBodyFormat = HttpError{
	Code:        "wrongBodyFormat",
	Status:      http.StatusBadRequest,
	Description: "bad request, body should be in json format.",
}
var httpError_CouldNotParseUser = HttpError{
	Code:        "couldNotParseUser",
	Status:      http.StatusBadRequest,
	Description: "bad request, Body could not be parsed to a valid User format.",
}
