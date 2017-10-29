package api

import "testing"

func TestNullRegistererGET(t *testing.T) {
	h := &handlerMock{}
	n := &NullRegisterer{}
	n.GET("prefix", "xxx", h)
}

func TestNullRegistererPATCH(t *testing.T) {
	h := &handlerMock{}
	n := &NullRegisterer{}
	n.GET("prefix", "xxx", h)
}

func TestNullRegistererPOST(t *testing.T) {
	h := &handlerMock{}
	n := &NullRegisterer{}
	n.GET("prefix", "xxx", h)
}

func TestNullRegistererPUT(t *testing.T) {
	h := &handlerMock{}
	n := &NullRegisterer{}
	n.GET("prefix", "xxx", h)
}

func TestNullRegistererDELETE(t *testing.T) {
	h := &handlerMock{}
	n := &NullRegisterer{}
	n.GET("prefix", "xxx", h)
}
