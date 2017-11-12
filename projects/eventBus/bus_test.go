package eventBus

import (
	"testing"
	"github.com/open-gtd/server/eventBus"
	"github.com/stretchr/testify/mock"
)

func TestInitializeWillSubscribeToTagCreated(t *testing.T) {
	bus := &busMock{}
	getBus = func() eventBus.Bus {
		return bus
	}

	bus.On("Subscribe", eventBus.Topic("tagCreated"), mock.Anything).Return(nil)

	Initialize()

	bus.AssertExpectations(t)
}
