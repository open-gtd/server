package tags

import "testing"

func TestInitializeShouldInitializeApiAndSse(t *testing.T) {

	initializerMock := &initializerMock{}

	apiInitializer = initializerMock.apiInitialize
	sseInitializer = initializerMock.sseInitialize

	initializerMock.On("apiInitialize")
	initializerMock.On("sseInitialize")

	Initialize()

	initializerMock.AssertExpectations(t)
}
