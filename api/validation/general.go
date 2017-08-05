package validation

import "github.com/open-gtd/server/api/validation/errors"

//AllowedValue test if value is allowed (within given list of allowed values, otherwise returns Not Allowed Value Validation Error
func AllowedValue(name string, value string, allowedValues []string) error {

	for _, allowedValue := range allowedValues {
		if allowedValue == value {
			return nil
		}
	}

	return errors.NewNotAllowedValue(name, value, allowedValues)
}
