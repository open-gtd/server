package errors

//ValidationError is as general type for Validation errors, when thrown causes BadRequest instead of InternalServerError
type ValidationError error

//IsValidationError returns true if err is of type ValidationError, otherwise false
func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}
