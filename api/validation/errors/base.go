package errors

//ValidationError - General type for Validation errors, When thrown causes BadRequest instead of InternalServerError
type ValidationError error
