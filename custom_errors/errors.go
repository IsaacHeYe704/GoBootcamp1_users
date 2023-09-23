package custom_errors

import "errors"

var Error_UserNotFound = errors.New("User not found")
var Error_UuidAlreadyExists = errors.New("there is already an user with this uui")
