package api

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage(t *testing.T) {

	const message = "ERROR message~"
	assert.Equal(t, message, Message(errors.New(message)).Message)

}
