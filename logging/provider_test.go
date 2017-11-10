package logging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetLogger(t *testing.T) {
	testSet(t, &NullLogger{})
	testSet(t, &loggerMock{})
}
func testSet(t *testing.T, mock Logger) {
	SetLogger(mock)
	assert.Equal(t, mock, l)
}

func TestGetLogger(t *testing.T) {
	testGet(t, &NullLogger{})
	testGet(t, &loggerMock{})
}

func testGet(t *testing.T, log Logger) {
	l = log
	result := GetLogger()
	assert.Equal(t, log, result)
}
