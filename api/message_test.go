package api

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {

	const message = "ERROR message~"
	assert.Equal(t, message, Message(errors.New(message)).Message)

}