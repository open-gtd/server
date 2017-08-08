package echo

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
	"github.com/stretchr/testify/mock"
)

type register struct {
	mock.Mock
}

func (r register) Add(method string, path string, h echo.HandlerFunc) {
	r.Called(method, path, h)
}

func TestRegisterer_GET_ShouldCallRegisterAddWithGETMethodAndParameters(t *testing.T) {
	groups := map[string]Router{}

	path := "/xcx"

	r := prepareRouterMock("GET", path)

	sut := NewEchoRegisterer(r, groups)
	sut.GET("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_POST_ShouldCallRegisterAddWithPOSTMethodAndParameters(t *testing.T) {
	groups := map[string]Router{}

	path := "/xzx"
	r := prepareRouterMock("POST", path)

	sut := NewEchoRegisterer(r, groups)
	sut.POST("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PATCH_ShouldCallRegisterAddWithPATCHMethodAndParameters(t *testing.T) {
	groups := map[string]Router{}

	path := "/xxz"
	r := prepareRouterMock("PATCH", path)

	sut := NewEchoRegisterer(r, groups)
	sut.PATCH("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PUT_ShouldCallRegisterAddWithPUTMethodAndParameters(t *testing.T) {
	groups := map[string]Router{}

	path := "/zxx"
	r := prepareRouterMock("PUT", path)

	sut := NewEchoRegisterer(r, groups)
	sut.PUT("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_DELETE_ShouldCallRegisterAddWithDELETEMethodAndParameters(t *testing.T) {
	groups := map[string]Router{}

	path := "/xyy"
	r := prepareRouterMock("DELETE", path)

	sut := NewEchoRegisterer(r, groups)
	sut.DELETE("", path, handler)

	r.AssertExpectations(t)
}

func prepareRouterMock(method string, path string) register {
	r := register{}

	r.On("Add", method, path, mock.AnythingOfType("echo.HandlerFunc"))
	return r
}

var handler = func(api.Request, api.Response) error {
	return nil
}
