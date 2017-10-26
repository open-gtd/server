package auth

import "testing"

func TestInitialize_ShouldCallApiInitialize(t *testing.T) {
	apiInitializeMock := &apiInitializeMock{}
	apiInitializer = apiInitializeMock.Initialize

	apiInitializeMock.On("Initialize")

	Initialize()

	apiInitializeMock.AssertExpectations(t)
}
