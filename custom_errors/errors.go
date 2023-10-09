package custom_errors

import (
	"errors"
	"fmt"
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
	WrongBodyFormat = "WrongBodyFormat"
)

//sentinel errors

var Error_UserNotFound = errors.New("user not found")
var Error_UuidAlreadyExists = errors.New("there is already an user with this uui")

var Error_WrongBodyFormat = errors.New("wrong body format")
var Error_ParsingJson = errors.New("could not parse Json to User")
