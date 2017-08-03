package validation

import "github.com/open-gtd/server/api/validation/errors"

//AllowedValue - Based on list, validates that Given Value is allowed
func AllowedValue(name string, value string, allowedValues []string) error {

	for _, allowedValue := range allowedValues {
		if allowedValue == value {
			return nil
		}
	}

	return errors.NewNotAllowedValue(name, value, allowedValues)
}
