package presenters

import (
	"net/http"
	"testing"

	"errors"
	"fmt"

	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/auth/presentation"
	"github.com/open-gtd/server/eventBus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_ShowAuthFailed_WillCallResponseNoContentWithForbiddenCode(t *testing.T) {
	response := &testResponse{}
	bus := &testBus{}

	response.On("NoContent", http.StatusForbidden).
		Return(nil)

	sut := NewLogin(response, bus)
	sut.ShowAuthFailed()

	response.AssertExpectations(t)
}

func TestLogin_ShowAuthFailed_ShouldReturnErorr_IfResponseNoContentReturnsError(t *testing.T) {
	const responseError = "ResponseError"

	response := &testResponse{}
	bus := &testBus{}

	response.On("NoContent", mock.Anything).
		Return(errors.New(responseError))

	sut := NewLogin(response, bus)
	err := sut.ShowAuthFailed()

	assert.EqualError(t, err, responseError)
}

func TestLogin_ShowAuthFailed_ShouldReturnNil_IfResponseNoContentReturnsNil(t *testing.T) {
	response := &testResponse{}
	bus := &testBus{}

	response.On("NoContent", mock.Anything).
		Return(nil)

	sut := NewLogin(response, bus)
	err := sut.ShowAuthFailed()

	assert.Nil(t, err)
}

func TestLogin_Show_ShouldPassArgumentToConvertFunction(t *testing.T) {
	const token = "toke3n"

	cm := &convertMock{}
	convertFunc = cm.Convert

	bus := &testBus{}
	bus.On("Publish", mock.Anything, mock.Anything)

	response := &testResponse{}
	response.On("JSON", mock.Anything, mock.Anything).Return(nil)

	cert := domain.Cert{Token: domain.Token(token)}

	cm.On("Convert", cert).Return(presentation.Cert{Token: token})
	sut := NewLogin(response, bus)

	sut.Show(cert)

	cm.AssertExpectations(t)
}

func TestLogin_Show_ShouldPublishArgumentToLeggedInTopicOnBus(t *testing.T) {
	const token = "toke3n"

	cm := &convertMock{}
	convertFunc = cm.Convert

	cert := domain.Cert{Token: domain.Token(token)}

	bus := &testBus{}
	bus.On("Publish", eventBus.Topic("loggedIn"), cert)

	response := &testResponse{}
	response.On("JSON", mock.Anything, mock.Anything).Return(nil)

	cm.On("Convert", mock.Anything).Return(presentation.Cert{Token: token})
	sut := NewLogin(response, bus)

	sut.Show(cert)

	bus.AssertExpectations(t)
}

func TestLogin_Show_ShouldSetResponseToOkWithAuthData(t *testing.T) {
	const token = "toke3n"

	cm := &convertMock{}
	convertFunc = cm.Convert

	cert := domain.Cert{Token: domain.Token(token)}

	bus := &testBus{}
	bus.On("Publish", mock.Anything, mock.Anything)

	authData := presentation.Cert{Token: token}

	response := &testResponse{}
	response.On("JSON", http.StatusOK, authData).Return(nil)

	cm.On("Convert", mock.Anything).Return(authData)
	sut := NewLogin(response, bus)

	sut.Show(cert)

	response.AssertExpectations(t)
}

type testResponse struct {
	mock.Mock
}

func (t *testResponse) String(code int, s string) error {
	args := t.Called(code, s)
	return args.Error(0)
}
func (t *testResponse) JSON(code int, i interface{}) error {
	args := t.Called(code, i)
	return args.Error(0)
}
func (t *testResponse) NoContent(code int) error {
	args := t.Called(code)
	return args.Error(0)
}

type testBus struct {
	mock.Mock
}

func (b *testBus) Subscribe(topic eventBus.Topic, fn eventBus.BusHandler) error {
	args := b.Called(topic)
	return args.Error(0)
}
func (b *testBus) Unsubscribe(topic eventBus.Topic, handler eventBus.BusHandler) error {
	args := b.Called(topic)
	return args.Error(0)
}
func (b *testBus) Publish(topic eventBus.Topic, arg interface{}) {
	b.Called(topic, arg)
}

type convertMock struct {
	mock.Mock
}

func (c *convertMock) Convert(cert domain.Cert) presentation.Cert {
	args := c.Called(cert)

	result := args.Get(0)
	rCert, ok := result.(presentation.Cert)
	if !ok {
		panic(fmt.Sprintf("Return failed because object wasn't coddect type: {%v}", result))
	}
	return rCert
}
