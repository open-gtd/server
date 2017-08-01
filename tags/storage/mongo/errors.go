package mongo

import "strings"

const notFoundError = "not found"
const duplicateKeyError = "duplicate key error index"

func isNotFound(err error) bool {
	return err.Error() == notFoundError
}

func isDuplicateKey(err error) bool {
	return strings.Contains(err.Error(), duplicateKeyError)
}
