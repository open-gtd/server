package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetReader(t *testing.T) {
	mock := &readerMock{}
	SetReader(mock)
	assert.Equal(t, mock, r)
}

func TestGetReader(t *testing.T) {
	mock := &readerMock{}
	r = mock
	result := GetReader()
	assert.Equal(t, result, mock)
}
