package errors

const (
	NotUniqueErrorMessage = "not unique"
)

type NotUniqueError struct {
}

func (e NotUniqueError) Error() string {
	return NotUniqueErrorMessage
}

func NewNotUnique() error {
	return NotUniqueError{}
}
