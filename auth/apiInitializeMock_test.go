package auth

import "github.com/stretchr/testify/mock"

type apiInitializeMock struct {
	mock.Mock
}

func (a *apiInitializeMock) Initialize() {
	a.Called()
}
