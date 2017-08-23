package errors

const (
	Validation ErrorType = iota
)

type ErrorType int

//TypedError is as general type for Validation errors, when thrown causes BadRequest instead of InternalServerError
type TypedError interface {
	error
	Type() ErrorType
}

//IsValidationError returns true if err is of type ValidationError, otherwise false
func IsValidationError(err error) bool {
	typedError, ok := err.(TypedError)
	if !ok {
		return false
	}

	return typedError.Type() == Validation
}
