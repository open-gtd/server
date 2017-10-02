package echo

import (
	"testing"

	"fmt"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
	"github.com/stretchr/testify/mock"
)

func TestRegisterer_GET_ShouldCallRegisterAddWithGETMethodAndParameters(t *testing.T) {
	path := "/xcx"

	groups := map[string]Router{}
	r := prepareRouterMock("GET", path)

	sut := NewEchoRegisterer(r, groups)
	sut.GET("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_POST_ShouldCallRegisterAddWithPOSTMethodAndParameters(t *testing.T) {
	path := "/xzx"

	groups := map[string]Router{}
	r := prepareRouterMock("POST", path)

	sut := NewEchoRegisterer(r, groups)
	sut.POST("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PATCH_ShouldCallRegisterAddWithPATCHMethodAndParameters(t *testing.T) {
	path := "/xxz"

	groups := map[string]Router{}
	r := prepareRouterMock("PATCH", path)

	sut := NewEchoRegisterer(r, groups)
	sut.PATCH("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PUT_ShouldCallRegisterAddWithPUTMethodAndParameters(t *testing.T) {
	path := "/zxx"

	groups := map[string]Router{}
	r := prepareRouterMock("PUT", path)

	sut := NewEchoRegisterer(r, groups)
	sut.PUT("", path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_DELETE_ShouldCallRegisterAddWithDELETEMethodAndParameters(t *testing.T) {
	prefix := ""
	path := "/xyy"

	groups := map[string]Router{}
	r := prepareRouterMock("DELETE", path)

	sut := NewEchoRegisterer(r, groups)
	sut.DELETE(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_GET_ShouldCallRegisterAddWithGETMethodAndParametersAndPrefix(t *testing.T) {
	path := "/xcx"
	prefix := "/prefix"

	groups := map[string]Router{}
	r := prepareRouterMock("GET", prefix+path)

	sut := NewEchoRegisterer(r, groups)
	sut.GET(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_POST_ShouldCallRegisterAddWithPOSTMethodAndParametersAndPrefix(t *testing.T) {
	path := "/xzx"
	prefix := "/prefix"

	groups := map[string]Router{}
	r := prepareRouterMock("POST", prefix+path)

	sut := NewEchoRegisterer(r, groups)
	sut.POST(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PATCH_ShouldCallRegisterAddWithPATCHMethodAndParametersAndPrefix(t *testing.T) {
	path := "/xxz"
	prefix := "/prefix"

	groups := map[string]Router{}
	r := prepareRouterMock("PATCH", prefix+path)

	sut := NewEchoRegisterer(r, groups)
	sut.PATCH(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PUT_ShouldCallRegisterAddWithPUTMethodAndParametersAndPrefix(t *testing.T) {
	path := "/zxx"
	prefix := "/prefix"

	groups := map[string]Router{}
	r := prepareRouterMock("PUT", prefix+path)

	sut := NewEchoRegisterer(r, groups)
	sut.PUT(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_DELETE_ShouldCallRegisterAddWithDELETEMethodAndParametersAndPrefix(t *testing.T) {
	prefix := "/prefix"
	path := "/xyy"

	groups := map[string]Router{}
	r := prepareRouterMock("DELETE", prefix+path)

	sut := NewEchoRegisterer(r, groups)
	sut.DELETE(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_GET_ShouldCallRegisterAddOnGroupWithGETMethodAndParameters(t *testing.T) {
	path := "/xxz"
	prefix := "/prefix"

	r := registerMock{}
	groups := map[string]Router{}
	groups["/prefix"] = prepareRouterMock("GET", path)

	sut := NewEchoRegisterer(r, groups)
	sut.GET(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_POST_ShouldCallRegisterAddOnGroupWithPOSTMethodAndParameters(t *testing.T) {
	path := "/xxz"
	prefix := "/prefix"

	r := registerMock{}
	groups := map[string]Router{}
	groups["/prefix"] = prepareRouterMock("POST", path)

	sut := NewEchoRegisterer(r, groups)
	sut.POST(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PATCH_ShouldCallRegisterAddOnGroupWithPATCHMethodAndParameters(t *testing.T) {
	path := "/xxz"
	prefix := "/prefix"

	r := registerMock{}
	groups := map[string]Router{}
	groups["/prefix"] = prepareRouterMock("PATCH", path)

	sut := NewEchoRegisterer(r, groups)
	sut.PATCH(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_PUT_ShouldCallRegisterAddOnGroupWithPUTMethodAndParameters(t *testing.T) {
	path := "/xxz"
	prefix := "/prefix"

	r := registerMock{}
	groups := map[string]Router{}
	groups["/prefix"] = prepareRouterMock("PUT", path)

	sut := NewEchoRegisterer(r, groups)
	sut.PUT(prefix, path, handler)

	r.AssertExpectations(t)
}

func TestRegisterer_DELETE_ShouldCallRegisterAddOnGroupWithDELETEMethodAndParameters(t *testing.T) {
	path := "/xxz"
	prefix := "/prefix"

	r := registerMock{}
	groups := map[string]Router{}
	groups["/prefix"] = prepareRouterMock("DELETE", path)

	sut := NewEchoRegisterer(r, groups)
	sut.DELETE(prefix, path, handler)

	r.AssertExpectations(t)
}

func prepareRouterMock(method string, path string) registerMock {
	r := registerMock{}

	r.On("Add", method, path, mock.AnythingOfType("echo.HandlerFunc")).Return(nil)
	return r
}

type testHandler struct{}

func (testHandler) Handle(request api.Request, response api.Response) error {
	return nil
}

var handler = testHandler{}

func route(obj interface{}) *echo.Route {
	var r *echo.Route
	var ok bool
	if obj == nil {
		return nil
	}
	if r, ok = obj.(*echo.Route); !ok {
		panic(fmt.Sprintf("assert: arguments: Controller failed because object wasn't correct type: %v", obj))
	}
	return r
}
