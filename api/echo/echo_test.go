package echo

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/labstack/echo"
)


const POST = "POST"
const GET = "GET"
const PATCH = "PATCH"
const PUT = "PUT"
const DELETE = "DELETE"

const emptyPrefix = ""

func TestRegisterer_GET_ShouldCallRegisterAddWithGETMethodAndParameters(t *testing.T) {
	const path = "/xcx"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(GET, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.GET(emptyPrefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_GET_RegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"

	routerMock := prepareRouterMockWithHandlerCall(GET, path)
	handlerFuncMock := &handlerFuncMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	groups := map[string]Router{}
	sut := NewEchoRegisterer(routerMock, groups)
	sut.GET(emptyPrefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_POST_ShouldCallRegisterAddWithPOSTMethodAndParameters(t *testing.T) {
	path := "/xzx"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(POST, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.POST(emptyPrefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_POST_RegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"

	routerMock := prepareRouterMockWithHandlerCall(POST, path)
	handlerFuncMock := &handlerFuncMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	groups := map[string]Router{}
	sut := NewEchoRegisterer(routerMock, groups)

	sut.POST(emptyPrefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_PATCH_ShouldCallRegisterAddWithPATCHMethodAndParameters(t *testing.T) {
	const path = "/xxz"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(PATCH, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.PATCH(emptyPrefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_PATCH_RegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"

	routerMock := prepareRouterMockWithHandlerCall(PATCH, path)
	handlerFuncMock := &handlerFuncMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	groups := map[string]Router{}
	sut := NewEchoRegisterer(routerMock, groups)
	sut.PATCH(emptyPrefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_PUT_ShouldCallRegisterAddWithPUTMethodAndParameters(t *testing.T) {
	const path = "/zxx"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(PUT, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.PUT(emptyPrefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_PUT_RegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"

	routerMock := prepareRouterMockWithHandlerCall(PUT, path)
	handlerFuncMock := &handlerFuncMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	groups := map[string]Router{}
	sut := NewEchoRegisterer(routerMock, groups)
	sut.PUT(emptyPrefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_DELETE_ShouldCallRegisterAddWithDELETEMethodAndParameters(t *testing.T) {
	const prefix = emptyPrefix
	const path = "/xyy"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(DELETE, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.DELETE(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_DELETE_RegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"

	routerMock := prepareRouterMockWithHandlerCall(DELETE, path)
	handlerFuncMock := &handlerFuncMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	groups := map[string]Router{}
	sut := NewEchoRegisterer(routerMock, groups)
	sut.DELETE(emptyPrefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_GET_ShouldCallRegisterAddWithGETMethodAndParametersAndPrefix(t *testing.T) {
	const path = "/xcx"
	const prefix = "/prefix"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(GET, prefix + path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.GET(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_POST_ShouldCallRegisterAddWithPOSTMethodAndParametersAndPrefix(t *testing.T) {
	const path = "/xzx"
	const prefix = "/prefix"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(POST, prefix + path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.POST(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_PATCH_ShouldCallRegisterAddWithPATCHMethodAndParametersAndPrefix(t *testing.T) {
	const path = "/xxz"
	const prefix = "/prefix"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(PATCH, prefix + path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.PATCH(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_PUT_ShouldCallRegisterAddWithPUTMethodAndParametersAndPrefix(t *testing.T) {
	const path = "/zxx"
	const prefix = "/prefix"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(PUT, prefix + path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.PUT(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_DELETE_ShouldCallRegisterAddWithDELETEMethodAndParametersAndPrefix(t *testing.T) {
	const prefix = "/prefix"
	const path = "/xyy"

	groups := map[string]Router{}

	handlerFuncMock := &handlerFuncMock{}
	routerMock := prepareRouterMock(DELETE, prefix + path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.DELETE(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_GET_ShouldCallRegisterAddOnGroupWithGETMethodAndParameters(t *testing.T) {
	const path = "/xxz"
	const prefix = "/prefix"

	handlerFuncMock := &handlerFuncMock{}
	routerMock := registerMock{}

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMock(GET, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.GET(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_GET_OnGroupRegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"
	const prefix = "/prefix"

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMockWithHandlerCall(GET, path)

	handlerFuncMock := &handlerFuncMock{}
	registerMock := registerMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	sut := NewEchoRegisterer(registerMock, groups)
	sut.GET(prefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_POST_ShouldCallRegisterAddOnGroupWithPOSTMethodAndParameters(t *testing.T) {
	const path = "/xxz"
	const prefix = "/prefix"

	routerMock := registerMock{}
	handlerFuncMock := &handlerFuncMock{}

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMock(POST, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.POST(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_POST_OnGroupRegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"
	const prefix = "/prefix"

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMockWithHandlerCall(POST, path)

	handlerFuncMock := &handlerFuncMock{}
	registerMock := registerMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	sut := NewEchoRegisterer(registerMock, groups)
	sut.POST(prefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_PATCH_ShouldCallRegisterAddOnGroupWithPATCHMethodAndParameters(t *testing.T) {
	const path = "/xxz"
	const prefix = "/prefix"

	routerMock := registerMock{}
	handlerFuncMock := &handlerFuncMock{}

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMock(PATCH, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.PATCH(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_PATCH_OnGroupRegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"
	const prefix = "/prefix"

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMockWithHandlerCall(PATCH, path)

	handlerFuncMock := &handlerFuncMock{}
	registerMock := registerMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	sut := NewEchoRegisterer(registerMock, groups)
	sut.PATCH(prefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_PUT_ShouldCallRegisterAddOnGroupWithPUTMethodAndParameters(t *testing.T) {
	const path = "/xxz"
	const prefix = "/prefix"

	routerMock := registerMock{}
	handlerFuncMock := &handlerFuncMock{}

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMock(PUT, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.PUT(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_PUT_OnGroupRegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"
	const prefix = "/prefix"

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMockWithHandlerCall(PUT, path)

	handlerFuncMock := &handlerFuncMock{}
	registerMock := registerMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	sut := NewEchoRegisterer(registerMock, groups)
	sut.PUT(prefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func TestRegisterer_DELETE_ShouldCallRegisterAddOnGroupWithDELETEMethodAndParameters(t *testing.T) {
	const path = "/xxz"
	const prefix = "/prefix"

	routerMock := registerMock{}
	handlerFuncMock := &handlerFuncMock{}

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMock(DELETE, path)

	sut := NewEchoRegisterer(routerMock, groups)
	sut.DELETE(prefix, path, handlerFuncMock)

	routerMock.AssertExpectations(t)
}

func TestRegisterer_DELETE_OnGroupRegisteredHandlerFuncShouldBeCalledWHenRouterCallRegisteredHandler(t *testing.T) {
	const path = "/xzx"
	const prefix = "/prefix"

	groups := map[string]Router{}
	groups[prefix] = prepareRouterMockWithHandlerCall(DELETE, path)

	handlerFuncMock := &handlerFuncMock{}
	registerMock := registerMock{}

	handlerFuncMock.On("Handle", mock.Anything, mock.Anything).
		Return(nil)

	sut := NewEchoRegisterer(registerMock, groups)
	sut.DELETE(prefix, path, handlerFuncMock)

	handlerFuncMock.AssertExpectations(t)
}

func prepareRouterMock(method string, path string) registerMock {
	registermock := registerMock{}

	var route *echo.Route
	registermock.On("Add", method, path, mock.AnythingOfType("echo.HandlerFunc")).Return(route)
	return registermock
}

func prepareRouterMockWithHandlerCall(method string, path string) registerMock {
	registerMock := registerMock{}

	var route *echo.Route
	registerMock.On("Add", method, path, mock.AnythingOfType("echo.HandlerFunc")).
		Run(func(args mock.Arguments){
			args.Get(2).(echo.HandlerFunc)(&contextMock{})
		}).
		Return(route)

	return registerMock
}
