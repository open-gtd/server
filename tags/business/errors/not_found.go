package errors

const (
	NotFoundErrorMessage = "not found"
)

type NotFoundError struct {
}

func (e NotFoundError) Error() string {
	return NotFoundErrorMessage
}

func NewNotFound() error {
	return NotFoundError{}
}
