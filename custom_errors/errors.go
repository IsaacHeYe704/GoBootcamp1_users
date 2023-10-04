package custom_errors

import "errors"

type HttpError struct {
	Code        string
	Status      int
	Description string
}
type ServiceError struct {
	Code        string
	Description string
}

var Error_UserNotFound = errors.New("user not found")
var Error_UuidAlreadyExists = errors.New("there is already an user with this uui")

var Error_WrongBodyFormat = errors.New("wrong body format")
var Error_ParsingJson = errors.New("could not parse Json to User")
