package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetRegisterer(t *testing.T) {
	registerer := TestRegisterer{}
	SetRegisterer(registerer)
	assert.Equal(t, registerer, r)
}

func TestGetRegisterer(t *testing.T) {
	registerer := TestRegisterer{}
	r = registerer

	assert.Equal(t, registerer, GetRegisterer())
}

type TestRegisterer struct {
}

func (TestRegisterer) GET(prefix string, path string, handlerFunc HandlerFunc)    {}
func (TestRegisterer) PATCH(prefix string, path string, handlerFunc HandlerFunc)  {}
func (TestRegisterer) POST(prefix string, path string, handlerFunc HandlerFunc)   {}
func (TestRegisterer) PUT(prefix string, path string, handlerFunc HandlerFunc)    {}
func (TestRegisterer) DELETE(prefix string, path string, handlerFunc HandlerFunc) {}
