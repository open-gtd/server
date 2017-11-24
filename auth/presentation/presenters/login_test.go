package presenters

import (
	"net/http"
	"testing"

	"errors"

	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/auth/presentation"
	"github.com/open-gtd/server/eventBus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_ShowAuthFailed_WillCallResponseNoContentWithForbiddenCode(t *testing.T) {
	response := &responseMock{}
	bus := &busMock{}

	response.On("NoContent", http.StatusForbidden).
		Return(nil)

	sut := NewLogin(response, bus)
	sut.ShowAuthFailed()

	response.AssertExpectations(t)
}

func TestLogin_ShowAuthFailed_ShouldReturnError_IfResponseNoContentReturnsError(t *testing.T) {
	const responseError = "ResponseError"

	response := &responseMock{}
	bus := &busMock{}

	response.On("NoContent", mock.Anything).
		Return(errors.New(responseError))

	sut := NewLogin(response, bus)
	err := sut.ShowAuthFailed()

	assert.EqualError(t, err, responseError)
}

func TestLogin_ShowAuthFailed_ShouldReturnNil_IfResponseNoContentReturnsNil(t *testing.T) {
	response := &responseMock{}
	bus := &busMock{}

	response.On("NoContent", mock.Anything).
		Return(nil)

	sut := NewLogin(response, bus)
	err := sut.ShowAuthFailed()

	assert.Nil(t, err)
}

func TestLogin_Show_ShouldPassArgumentToConvertFunction(t *testing.T) {
	const token = "toke3n"

	cm := &convertFuncMock{}
	convertFunc = cm.Convert

	bus := &busMock{}
	bus.On("Publish", mock.Anything, mock.Anything)

	response := &responseMock{}
	response.On("JSON", mock.Anything, mock.Anything).Return(nil)

	cert := domain.Cert{Token: domain.Token(token)}

	cm.On("Convert", cert).Return(presentation.Cert{Token: token})
	sut := NewLogin(response, bus)

	sut.Show(cert)

	cm.AssertExpectations(t)
}

func TestLogin_Show_ShouldPublishArgumentToLeggedInTopicOnBus(t *testing.T) {
	const token = "toke3n"

	cm := &convertFuncMock{}
	convertFunc = cm.Convert

	cert := domain.Cert{Token: domain.Token(token)}

	bus := &busMock{}
	bus.On("Publish", eventBus.Topic("loggedIn"), cert)

	response := &responseMock{}
	response.On("JSON", mock.Anything, mock.Anything).Return(nil)

	cm.On("Convert", mock.Anything).Return(presentation.Cert{Token: token})
	sut := NewLogin(response, bus)

	sut.Show(cert)

	bus.AssertExpectations(t)
}

func TestLogin_Show_ShouldSetResponseToOkWithAuthData(t *testing.T) {
	const token = "toke3n"

	cm := &convertFuncMock{}
	convertFunc = cm.Convert

	cert := domain.Cert{Token: domain.Token(token)}

	bus := &busMock{}
	bus.On("Publish", mock.Anything, mock.Anything)

	authData := presentation.Cert{Token: token}

	response := &responseMock{}
	response.On("JSON", http.StatusOK, authData).Return(nil)

	cm.On("Convert", mock.Anything).Return(authData)
	sut := NewLogin(response, bus)

	sut.Show(cert)

	response.AssertExpectations(t)
}
